
ALL := $(patsubst %.go,%,$(wildcard *.go))

all: $(ALL)

$(ALL): $@
	go build -o $@ $@.go

clean:
	rm -f $(ALL)

.PHONY: all clean
