ORG=cipherutils
PREFIX?=/usr/local
DIRS=bin share
INSTALL_DIRS=`find $(DIRS) -type d`
INSTALL_FILES=`find $(DIRS) -type f`
DOCS=$(shell find docs/*.1.md -type f | sed 's/docs/share\/man\/man1/g' | sed 's/\.md//g')
TOOLS=$(shell for tool in `find */main.go -type f`; do dirname $$tool; done)
PROGRAMS=$(shell for tool in `find */main.go -type f`; do echo bin/`dirname $$tool`; done)
VERSION?=$(shell git describe --abbrev=0 --tags)

all: dictionary/words.go $(PROGRAMS) $(DOCS)
.PHONY: all

bin/%:
	@mkdir -p bin
	@cd $(@F) && go mod download && go build -o ../bin/$(@F)

dictionary/words.go:
	@cd dictionary && ruby goify.rb

install: $(PROGRAMS)
	@for dir in $(INSTALL_DIRS); do install $$dir $(PREFIX)/$$dir; done
.PHONY: install

uninstall:
	@for file in $(INSTALL_FILES); do rm -rf $$file; done
.PHONY: uninstall

share/man/man1:
	@mkdir -p share/man/man1
share/man/man1/%.1: share/man/man1
	@ronn --roff --date="2014-11-01" --manual="User Manual" --organization="$(ORG)" --style=dark docs/$(@F).md
	@mv docs/$(@F) share/man/man1/$(@F)

clean:
	@rm -rf bin share dist
.PHONY: clean

dist: $(PROGRAMS)
	@mkdir -p dist/cipherutils-$(VERSION)
	@for dir in $(INSTALL_DIRS); do cp -a $$dir dist/cipherutils-$(VERSION)/$$dir; done
	@cp README.md dist/cipherutils-$(VERSION)/README.md
	@cp Makefile dist/cipherutils-$(VERSION)/Makefile
	@cd dist; tar -czf cipherutils-$(VERSION).tar.gz ./cipherutils-$(VERSION)
.PHONY: dist

distclean:
	@rm -rf dist
.PHONY: distclean

mostlyclean:
	@rm -rf bin share
.PHONY: mostlyclean

maintainer-clean:
	@rm -rf share
.PHONY: maintainer-clean

check:
	@for tool in $(TOOLS); do pushd $$tool > /dev/null 2>&1; go test; popd > /dev/null 2>&1; done
.PHONY: check

serve:
	@jekyll serve -s docs
.PHONY: serve
