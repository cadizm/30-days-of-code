
all: hello datatype operators conditionals clsvsins

hello:
	go build -o ./bin/hello github.com/cadizm/30-days-of-code/hello

datatype:
	go build -o ./bin/datatype github.com/cadizm/30-days-of-code/datatype

operators:
	go build -o ./bin/operators github.com/cadizm/30-days-of-code/operators

conditionals:
	go build -o ./bin/conditionals github.com/cadizm/30-days-of-code/conditionals

clsvsins:
	go build -o ./bin/clsvsins github.com/cadizm/30-days-of-code/clsvsins

clean:
	@rm -f ./bin/*

.PHONY: clean hello datatype operators conditionals clsvsins
