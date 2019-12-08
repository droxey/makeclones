#
# Makefile
#
VERSION = snapshot
GHRFLAGS =
.PHONY: build release

default: build

build:
	goxc -d=pkg -pv=$(VERSION) -bc="darwin,linux"

release:
	ghr $(GHRFLAGS) v$(VERSION) pkg/$(VERSION)
