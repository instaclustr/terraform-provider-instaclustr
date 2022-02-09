BIN_NAME=terraform-provider-instaclustr

VERSION=2.0.0

INSTALL_FOLDER=$(HOME)/.terraform.d/plugins/terraform.instaclustr.com/instaclustr/instaclustr/$(VERSION)/darwin_amd64

.PHONY: local-build preprod-build build install local-gen-docs gen-docs

define create-build
	go build $(FLAGS) -ldflags="-X 'main.ProviderHost=$(1)'" -o bin/$(BIN_NAME)_v$(VERSION) main.go
endef

local-build:
	$(call create-build,http://localhost:8090)

preprod-build:
	$(call create-build,https://api.dev.instaclustr.com)

build:
	$(call create-build,https://api.instaclustr.com)

install:
	@if [ ! -d $(INSTALL_FOLDER) ]; then \
		echo "$(INSTALL_FOLDER) doesn't exist, creating..."; \
		mkdir -p $(INSTALL_FOLDER); \
	fi
	cp ./bin/$(BIN_NAME)_v$(VERSION) $(INSTALL_FOLDER)

local-gen-docs:
	IC_TF_PROVIDER_SWAGGER_HOST_OVERRIDE=http://localhost:8090 go generate

gen-docs:
	go generate
