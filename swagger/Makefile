interactive-docker:
	cidocker run -it --rm -v ${PWD}:/go/src -w /go/src -p 8185:8185 -p 42350:42350 golang:latest

swagger-install:
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger-validate:
	swagger validate ./URCapAPI.yaml

swagger-generate-model:
	swagger generate model \
		--spec ./URCapAPI.yaml \
		--keep-spec-order \
		--struct-tags yaml

swagger-generate-server:
	swagger generate server \
		--spec ./URCapAPI.yaml \
		--keep-spec-order \
		--api-package=operations \
		--model-package=models \
		--struct-tags yaml \
		--strict-responders \
		--strict-additional-properties

swagger-generate-cli:
	swagger generate cli \
	--spec ./URCapAPI.yaml \
	--keep-spec-order \
	--struct-tags yaml

swagger-server:
	swagger serve --port=42350 --no-open -F=swagger URCapAPI.yaml
