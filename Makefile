
all: hello

hello:
	go build -o ./bin/hello github.com/cadizm/30-days-of-code/hello

clean:
	@rm -f ./bin/*

.PHONY: clean hello
