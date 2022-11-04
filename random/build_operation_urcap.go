package cli

import (
	"archive/zip"
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"universal-robots/urcapctl/models"

	weaselDocker "github.com/codetent/weasel/pkg/weasel/docker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	cp "github.com/otiai10/copy"
)

func Build(urcapSRCPath, urcapDstPath string) (err error) {

	// load manifest
	manifest, err := ManifestLoad(path.Join(urcapSRCPath, "Manifest.yaml"))
	if err != nil {
		return err
	}

	// create a temp build directory
	buildPath, err := os.MkdirTemp("", "build")
	if err != nil {
		return nil
	}
	defer os.RemoveAll(buildPath)

	// build container artifact
	for _, config := range manifest.Artifacts.Containers {
		err = buildContainerArtifact(urcapSRCPath, buildPath, config)
		if err != nil {
			return
		}
	}

	// build polyscope artifact
	for _, config := range manifest.Artifacts.PolyscopeBundles {
		err = buildPolyscopeBundleArtifact(urcapSRCPath, buildPath, config)
		if err != nil {
			return
		}
	}

	// build webarchive artifact
	for _, config := range manifest.Artifacts.WebArchives {
		err = buildWebArchiveArtifact(urcapSRCPath, buildPath, config)
		if err != nil {
			return
		}
	}

	// copy manifest.yaml
	err = cp.Copy(path.Join(urcapSRCPath, "manifest.yaml"), path.Join(buildPath, "manifest.yaml"))
	if err != nil {
		return
	}

	// copy licenses
	licensePath := path.Join(urcapSRCPath, "LICENSE")
	_, err = os.Stat(licensePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("License file doesn't exsits")
		} else {
			return
		}
	} else {
		err = cp.Copy(licensePath, path.Join(buildPath, "LICENSE"))
		if err != nil {
			return err
		}
	}

	// zip the content
	urcapxPath := path.Join(urcapDstPath, manifest.Metadata.UrcapID+".urcapx")
	err = compressDir(buildPath, urcapxPath, false)
	if err != nil {
		return err
	}

	return nil
}

func compressDir(src, dst string, includeSrc bool) error {
	zipfile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		var name string
		if includeSrc {
			name, err = filepath.Rel(filepath.Dir(src), path)
			if err != nil {
				return err
			}
		} else {
			if path == src {
				return nil
			}
			name, err = filepath.Rel(src, path)
			if err != nil {
				return err
			}
		}

		name = filepath.ToSlash(name)
		if info.IsDir() {
			name += "/"
		}

		writer, err := archive.Create(name)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}

func buildWebArchiveArtifact(urcapSRCPath, buildPath string, config *models.WebArchive) (err error) {

	// get context path
	contextPath := path.Join(urcapSRCPath, *config.Name)
	fileInfo, err := os.Stat("contextPath")
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("docker context path: %s, must be a directory", contextPath)
	}

	// create dst directory
	dstPath := filepath.Join(buildPath, *config.Name)
	err = os.MkdirAll(dstPath, os.ModePerm)
	if err != nil {
		return
	}

	// move web-archive
	err = cp.Copy(contextPath, dstPath)
	if err != nil {
		return
	}

	return
}

func buildPolyscopeBundleArtifact(urcapSRCPath, buildPath string, config *models.PolyscopeBundle) (err error) {

	// get context path
	contextPath := path.Join(urcapSRCPath, *config.Name)
	fileInfo, err := os.Stat("contextPath")
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("docker context path: %s, must be a directory", contextPath)
	}

	// mvn command
	cmd := exec.Command("mvn", "package")
	cmd.Dir = contextPath
	err = cmd.Run()
	if err != nil {
		return err
	}

	// get path to built jar file
	var jars []string
	filepath.Walk(cmd.Dir, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() && filepath.Ext(path) == ".jar" {
			jars = append(jars, path)
		}
		return nil
	})
	bundle := jars[0]
	bundle = filepath.Base(bundle)

	// create target path
	dstPath := path.Join(buildPath, *config.Name)
	err = os.MkdirAll(dstPath, os.ModePerm)
	if err != nil {
		return err
	}

	// copy to dst
	err = cp.Copy(path.Join(contextPath, "target", bundle), path.Join(dstPath, bundle))
	if err != nil {
		return err
	}

	// clean
	cmd = exec.Command("mvn", "clean")
	cmd.Dir = contextPath
	err = cmd.Run()
	if err != nil {
		return err
	}

	return
}

func buildContainerArtifact(urcapSRCPath, buildPath string, config *models.Container) (err error) {

	// get context path
	contextPath := path.Join(urcapSRCPath, *config.Name)
	fileInfo, err := os.Stat("contextPath")
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("docker context path: %s, must be a directory", contextPath)
	}

	// build image
	err = buildImage(contextPath, *config.Image)
	if err != nil {
		return
	}

	// create target path
	dstPath := path.Join(buildPath, *config.Name)
	err = os.MkdirAll(dstPath, os.ModePerm)
	if err != nil {
		return err
	}

	// save image
	err = saveImage(dstPath, *config.Name)
	if err != nil {
		return
	}

	return
}

func buildImage(contextPath, tag string) (err error) {

	// check if Dockerfile exist in context path
	_, err = os.Stat(path.Join(contextPath, "Dockerfile"))
	if err != nil {
		return
	}

	// new docker client
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return
	}

	// create io.reader
	tarReader, err := archive.TarWithOptions(contextPath, &archive.TarOptions{})
	if err != nil {
		return
	}

	response, err := client.ImageBuild(context.Background(), tarReader, types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       []string{tag},
		Remove:     true,
	})
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if debug {
		log.Printf("res.Body is = %v", response.Body)
		scanner := bufio.NewScanner(response.Body)
		for scanner.Scan() {
			lastLine := scanner.Text()
			fmt.Println(lastLine)
		}
	}
	return nil
}
func saveImage(targetDirPath, tag string) (err error) {

	// get imageId
	imageId, err := weaselDocker.ImageIdByTag(tag)
	if err != nil {
		return
	}

	archivePath := filepath.Join(targetDirPath, imageId+".tgz")
	_, err = os.Stat(archivePath)
	if errors.Is(err, os.ErrNotExist) {
		err = weaselDocker.ImageExport(imageId, archivePath)
		if err != nil {
			return
		}
	}

	return
}
