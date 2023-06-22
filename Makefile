BIN_NAME=terraform-provider-instaclustr

VERSION=2.0.53

FULL_BIN_NAME="${BIN_NAME}_v${VERSION}"
SHASUM_NAME="${BIN_NAME}_${VERSION}_SHA256SUMS"

INSTALL_FOLDER=$(HOME)/.terraform.d/plugins/terraform.instaclustr.com/instaclustr/instaclustr/$(VERSION)/darwin_amd64

.PHONY: local-build preprod-build build build-all install local-gen-docs gen-docs release_version

release_version:
	@echo v$(VERSION)

build:
	go build $(FLAGS) -o bin/${FULL_BIN_NAME} main.go


build-all-platforms:
	rm -f bin/${BIN_NAME}*.zip
	rm -f bin/${SHASUM_NAME}
	env GOOS=darwin GOARCH=amd64 make build
	zip -j bin/${FULL_BIN_NAME}_darwin_amd64.zip bin/${FULL_BIN_NAME}
	env GOOS=darwin GOARCH=arm64 make build
	zip -j bin/${FULL_BIN_NAME}_darwin_arm64.zip bin/${FULL_BIN_NAME}
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make build FLAGS="-ldflags '-w -s -extldflags \"-static\"' -tags netgo -a -v"
	zip -j bin/${FULL_BIN_NAME}_linux_amd64.zip bin/${FULL_BIN_NAME}
	env GOOS=linux GOARCH=arm64 make build
	zip -j bin/${FULL_BIN_NAME}_linux_arm64.zip bin/${FULL_BIN_NAME}
	env GOOS=linux GOARCH=arm GOARM=6 make build
	zip -j bin/${FULL_BIN_NAME}_linux_arm.zip bin/${FULL_BIN_NAME}
	env GOOS=windows GOARCH=amd64 make build
	zip -j bin/${FULL_BIN_NAME}_windows_amd64.zip bin/${FULL_BIN_NAME}
	env GOOS=windows GOARCH=386 make build
	zip -j bin/${FULL_BIN_NAME}_windows_386.zip bin/${FULL_BIN_NAME}
	rm bin/${FULL_BIN_NAME} # remove lingering build
	shasum -a 256 bin/*.zip > bin/${SHASUM_NAME}
	sed -i '' -e 's/bin\///g' bin/${SHASUM_NAME}

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
