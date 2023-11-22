install_swagger:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger: install_swagger
	swagger generate spec -o ./swagger.yaml --scan-models

create_sdk_dir:
	mkdir -p sdk

generate_client: create_sdk_dir
	cd sdk && swagger generate client -f ../swagger.yaml -A product-api
