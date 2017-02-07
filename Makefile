
ALL := $(patsubst %.go,%,$(wildcard *.go))

all:
	@for prog in $(ALL); do \
		echo go build -o $$prog $$prog.go; \
		go build -o $$prog $$prog.go; \
	done;

$(ALL): $@
	go build -o $@ $@.go

clean:
	@for prog in $(ALL); do \
		rm -f $$prog; \
	done;

.PHONY: all clean $(ALL)
