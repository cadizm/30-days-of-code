
all: hello datatype operators

hello:
	go build -o ./bin/hello github.com/cadizm/30-days-of-code/hello

datatype:
	go build -o ./bin/datatype github.com/cadizm/30-days-of-code/datatype

operators:
	go build -o ./bin/operators github.com/cadizm/30-days-of-code/operators

clean:
	@rm -f ./bin/*

.PHONY: clean hello datatype operators
