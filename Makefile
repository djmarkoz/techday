VERBOSE_1 := -v
VERBOSE_2 := -v -x

.PHONY: all build install uninstall test run vendor clean cleanall docker docker-run help

WHAT := techday-cli
		
all: build test

build: vendor
	for target in $(WHAT); do \
		$(BUILD_ENV_FLAGS) go build $(VERBOSE_$(V)) -o bin/$$target  ./cmd/$$target; \
	done

install: vendor
	for target in $(WHAT); do \
		$(BUILD_ENV_FLAGS) go install $(VERBOSE_$(V)) ./cmd/$$target; \
	done

uninstall:
	for target in $(WHAT); do \
		rm $$GOPATH/bin/$$target; \
	done

test:
	go test ./pkg/... ./cmd/... ./internal/...
	
vet:
	go vet ./pkg/... ./cmd/... ./internal/...
		
run:
	./bin/techday-cli

vendor:
	dep ensure

clean:
	rm -rf ./bin

cleanall: clean
	rm -rf ./vendor
		
docker:
	docker build -t techday-cli .
	
docker-run:
	docker run --rm -ti techday-cli    

help:
	@echo "make environment variables"
	@echo "  V                 - Build verbosity {0,1,2}."
	@echo "  BUILD_ENV_FLAGS   - Environment added to 'go build'."
	@echo "  WHAT              - Command to build. (e.g. WHAT=techday-cli)"
