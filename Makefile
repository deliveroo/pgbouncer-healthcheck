APPNAME := pgbouncer-healthcheck
VERSION := $(shell cat VERSION)

src_files = *.go Gopkg.* VERSION
docker = Dockerfile build.sh vendor.sh
test_files = tests/* tests/scripts/*

build: $(APPNAME)

$(APPNAME): $(src_files) $(docker) $(test_files)
	-rm -fv $(APPNAME)
	docker build --target builder -t $(APPNAME) .
	docker run --rm $(APPNAME) tar c .|tar x ./$(APPNAME)
	# Docker will use cache if nothing has changed, meaning
	# the binary might have an old timestamp. Update to make
	# sure Make doesn't try to rebuild each time
	touch $(APPNAME)

test: $(APPNAME)
	tests/run_tests.sh $(APPNAME)

stash: $(APPNAME)
	mkdir -p workspace
	zip workspace/$(APPNAME)-$(CIRCLE_SHA1).zip $(APPNAME)

unstash:
	unzip workspace/$(APPNAME)-$(CIRCLE_SHA1).zip

upload:
	aws s3 cp "$(APPNAME)" "s3://roo-apps-private-binaries/$(APPNAME)-$(VERSION)"

clean:
	-rm -fv $(APPNAME)
	-rm -rfv workspace

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

.PHONY: build clean dep test stash unstash upload
