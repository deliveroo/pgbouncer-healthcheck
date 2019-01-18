APPNAME := pgbouncer-healthcheck
VERSION := $(shell cat VERSION)

build: $(APPNAME)-$(VERSION).tar.gz

$(APPNAME): *.go Gopkg.*
	-rm -fv $(APPNAME)
	docker build --target builder -t $(APPNAME) .
	docker run --rm $(APPNAME) tar c .|tar x ./$(APPNAME)

$(APPNAME)-$(VERSION).tar.gz: $(APPNAME)
	-rm -fv $(APPNAME)-$(VERSION).tar.gz
	tar cvzf $(APPNAME)-$(VERSION).tar.gz $(APPNAME)

test: $(APPNAME)
	tests/run_tests.sh $(APPNAME)

clean:
	-rm $(APPNAME)
	-rm $(APPNAME)-$(VERSION).tar.gz

# Use this target to run dep commands
#
# examples:
#
# Initialise dep
# $ make dep COMMAND='init -v'
#
# After adding new dependencies in code
# $ make dep COMMAND='ensure -v -update'
#
dep:
	docker run -it --rm -v $(shell pwd):/go/src/github.com/deliveroo/$(APPNAME) $(APPNAME) dep $(COMMAND)

.PHONY: build clean dep test
