BINARY := pptgrep
LDFLAGS := -ldflags="-s -w"

.PHONY: bin
bin: deps
	for os in darwin linux windows; do \
	  for arch in amd64 386; do \
	    GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$$os-$$arch/$(BINARY); \
	  done; \
	done

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: deps
deps:
	glide install

.PHONY: install
install:
	go install $(LDFLAGS)

.PHONY: test
test:
	go test -cover -v `glide novendor`
