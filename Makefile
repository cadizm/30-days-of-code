
all: hello datatype

hello:
	go build -o ./bin/hello github.com/cadizm/30-days-of-code/hello

datatype:
	go build -o ./bin/datatype github.com/cadizm/30-days-of-code/datatype

clean:
	@rm -f ./bin/*

.PHONY: clean hello datatype
