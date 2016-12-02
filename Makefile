BINARY := pptgrep
LDFLAGS := -ldflags="-s -w"

bin:	
	go build $(LDFLAGS)

clean:
	rm -f pptgrep
