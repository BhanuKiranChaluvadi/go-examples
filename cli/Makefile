validate-swagger:
	swagger validate ./swagger.yaml

serve-swagger:
	swagger serve --port=42350 --no-open -F=swagger swagger.yaml

old-generate-cli:
	swagger generate cli \
		--spec ./swagger.yaml \
		--keep-spec-order \
		--struct-tags yaml

generate-cli:
	swagger generate cli \
	--spec=swagger.yaml \
	--cli-app-name urcapctl 

clean:
	rm -rf cli client cmd models && rm urcapx swagger.yaml

build:
	go build -o cmd/urcapctl/urcapctl cmd/urcapctl/main.go

copy:
	cp ~/git/urcapx-sdk-api/swagger.yaml swagger.yaml

init:
	go mod init universal-robots/urcapctl


