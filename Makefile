VERBOSE_1 := -v
VERBOSE_2 := -v -x

.PHONY: all build test run vendor clean cleanall tools docker docker-run help

WHAT := techday-cli
		
all: test build

build: vendor
	for target in $(WHAT); do \
		$(BUILD_ENV_FLAGS) go build $(VERBOSE_$(V)) -o bin/$$target -ldflags "-X $(REPOPATH).Version=$(VERSION)" ./cmd/$$target; \
	done

test:
	go test ./...
	
vet:
	go vet ./...
	
fmt:
	goimports -l -w .

run: build
	./bin/techday-cli

vendor:
	dep ensure

clean:
	rm -rf ./bin

cleanall: clean
	rm -rf ./vendor

tools:
	go get golang.org/x/tools/cmd/goimports
	
docker:
	docker build -t techday-cli .
	
docker-run: docker
	docker run --rm -ti techday-cli    

help:
	@echo "make environment variables"
	@echo "  V                 - Build verbosity {0,1,2}."
	@echo "  BUILD_ENV_FLAGS   - Environment added to 'go build'."
	@echo "  WHAT              - Command to build. (e.g. WHAT=techday-cli)"
