GOSWAGGER_IMAGE=quay.io/goswagger/swagger
GOSWAGGER_VERSION=v0.27.0
SWAGGERCMD=docker run --rm -v $(HOME):$(HOME) -w $(CURDIR) $(GOSWAGGER_IMAGE):$(GOSWAGGER_VERSION)
SWAGGER_SPEC_FILE=swagger.yaml

$(SWAGGER_SPEC_FILE):
	# Use swagger spec from openHAB
	curl -o $(SWAGGER_SPEC_FILE) https://raw.githubusercontent.com/openhab/openhab-addons/main/bundles/org.openhab.binding.tado/src/main/api/tado-api.yaml

.PHONY: test
test:
	go test .

.PHONY: generate
generate: $(SWAGGER_SPEC_FILE)
	go generate ./...
