interactive-docker:
	cidocker run -it --rm \
	-v ${PWD}:/go/src \
	-w /go/src \
	-p 42350:42350 \
	golang:latest

swagger-install:
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger-validate:
	swagger validate ./URCapAPI.yaml

swagger-serve:
	swagger serve --port=42350 --no-open -F=swagger URCapAPI.yaml

swagger-flatten:
	swagger flatten --output=URCapAPI.yaml --with-flatten minimal --format=yaml simpleAPI.yaml
