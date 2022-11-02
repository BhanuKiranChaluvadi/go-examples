interactive-docker:
	cidocker run -it --rm -v ${PWD}:/go/src -w /go/src -p 8185:8185 -p 42350:42350 golang:latest

install-swagger:
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

validate-swagger:
	swagger validate ./simpleAPI.yaml

generate-model:
	swagger generate model \
		--spec ./simpleAPI.yaml \
		--keep-spec-order \
		--struct-tags yaml

generate-server:
	swagger generate server \
		--spec ./URCapAPI.yaml \
		--keep-spec-order \
		--api-package=operations \
		--model-package=models \
		--struct-tags yaml \
		--strict-responders \
		--strict-additional-properties

generate-cli:
	swagger generate cli \
	--spec ./URCapAPI.yaml \
	--keep-spec-order \
	--struct-tags yaml

server-swagger:
	swagger serve --port=42350 --no-open -F=swagger swagger.yaml

flatten-swagger:
	swagger flatten -o swagger.yaml --with-flatten minimal --format=yaml simpleAPI.yaml

expand-swagger:
	swagger expand --output=swagger.yaml --format=yaml simpleAPI.yaml