interactive-docker:
	cidocker run -it --rm -v ${PWD}:/go/src -w /go/src -p 8185:8185 -p 42350:42350 golang:latest

install-swagger:
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

validate-swagger:
	swagger validate ./URCapAPI.yaml

clean:
	rm -rf cmd models restapi

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

run:
	go run cmd/urcap-api-server/main.go --port 8185

own:
	sudo chown -R ${USER}:${USER} *

server-swagger:
	swagger serve --port=42350 --no-open -F=swagger URCapAPI.yaml
	
generate-another-server:
	swagger validate ./URCapAPI.yaml
	swagger generate server --spec ./URCapAPI.yaml -t ${GOPATH}/src/server
	swagger generate server --spec ./URCapAPI.yaml --name urservice -t ${GOPATH}/src/server
	swagger generate server \
		--spec ./URCapAPI.yaml
		--api-package op \
		--model-package model \
		--strict-responders \
		--strict-additional-properties \
		--keep-spec-order \
		--principal univeral-robots/urservice \
		--exclude-main 