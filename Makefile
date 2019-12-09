#
# Makefile
#
VERSION = snapshot
GHRFLAGS =
.PHONY: build release

default: build

build:
	goxc -d=pkg -pv=$(VERSION) -bc="darwin"

release:
	ghr  -u droxey  $(GHRFLAGS) v$(VERSION) pkg/$(VERSION)
