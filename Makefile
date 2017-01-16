BINARY := pptgrep
LDFLAGS := -ldflags="-s -w"

bin:
	GOOS=windows GOARCH=386 go build $(LDFLAGS) -o ./bin/windows386/$(BINARY).exe
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o ./bin/windows64/$(BINARY).exe
	GOOS=darwin GOARCH=386 go build $(LDFLAGS) -o ./bin/mac386/$(BINARY)
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o ./bin/mac64/$(BINARY)
	GOOS=linux GOARCH=386 go build $(LDFLAGS) -o ./bin/linux386/$(BINARY)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o ./bin/linux64/$(BINARY)
clean:
	rm -rf ./bin
