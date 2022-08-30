BIN_NAME=terraform-provider-instaclustr

VERSION=2.0.0

INSTALL_FOLDER=$(HOME)/.terraform.d/plugins/terraform.instaclustr.com/instaclustr/instaclustr/$(VERSION)/darwin_amd64

.PHONY: local-build preprod-build build install local-gen-docs gen-docs

build:
	go build $(FLAGS) -o bin/$(BIN_NAME)_v$(VERSION) main.go

install:
	@if [ ! -d $(INSTALL_FOLDER) ]; then \
		echo "$(INSTALL_FOLDER) doesn't exist, creating..."; \
		mkdir -p $(INSTALL_FOLDER); \
	fi
	cp ./bin/$(BIN_NAME)_v$(VERSION) $(INSTALL_FOLDER)

local-gen-docs:
	IC_API_URL=http://localhost:8090 ./scripts/instaclustr-terraform-registry-documentation-update.sh

gen-docs:
	./scripts/instaclustr-terraform-registry-documentation-update.sh
