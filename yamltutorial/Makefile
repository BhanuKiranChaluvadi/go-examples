clean:
	rm -rf cmd models restapi

generate-model:
	swagger generate model \
		--spec ./simpleAPI.yaml \
		--struct-tags yaml

generate-server:
	swagger generate server \
		--spec ./simpleAPI.yaml \
		--api-package=operations \
		--model-package=models \
		--struct-tags=yaml \
		--strict-responders \
		--strict-additional-properties \
		--keep-spec-order