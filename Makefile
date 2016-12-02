BINARY := pptgrep
LDFLAGS := -ldflags="-s -w"

bin:	
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o ./bin/windows/$(BINARY).exe
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o ./bin/mac/$(BINARY)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o ./bin/linux/$(BINARY)
clean:
	rm -rf ./bin
