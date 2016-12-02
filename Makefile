BINARY := pptgrep
LDFLAGS := -ldflags="-s -w"

bin:	
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS)
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o pptgrep_x
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o pptgrep_l
clean:
	rm -f pptgrep*
