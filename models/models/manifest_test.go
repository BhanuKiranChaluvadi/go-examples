package models

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/go-openapi/strfmt"
	"gopkg.in/yaml.v2"
)

func TestManifest_Validate(t *testing.T) {
	type fields struct {
		APIVersion *string
		Artifacts  *ManifestArtifacts
		Metadata   *Metadata
	}
	type args struct {
		formats strfmt.Registry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Sample",
			fields: fields{
				APIVersion: new(string),
				Artifacts:  &ManifestArtifacts{},
				Metadata:   &Metadata{},
			},
			args: args{
				formats: strfmt.Default,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manifest{
				APIVersion: tt.fields.APIVersion,
				Artifacts:  tt.fields.Artifacts,
				Metadata:   tt.fields.Metadata,
			}
			if err := m.Validate(tt.args.formats); (err != nil) != tt.wantErr {
				t.Errorf("Manifest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManifest_Validate2(t *testing.T) {
	// read manifest
	buf, _ := ioutil.ReadFile("manifest_example.yaml")

	// input := []interface{}{}
	// json.Unmarshal(buf, input)

	// unmarshall into model
	model := Manifest{
		APIVersion: new(string),
		Artifacts:  &ManifestArtifacts{},
		Metadata:   &Metadata{},
	}

	// err := model.UnmarshalJSON(buf)

	err := yaml.Unmarshal(buf, model)
	if err != nil {
		fmt.Println("Error")
	}
}
