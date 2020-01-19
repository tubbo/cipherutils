ORG=cipherutils
PREFIX?=/usr/local
DIRS=bin share
INSTALL_DIRS=`find $(DIRS) -type d`
INSTALL_FILES=`find $(DIRS) -type f`
DOCS=$(shell find man/*.md -type f | sed 's/man/share\/man\/man1/g' | sed 's/\.md//g')
TOOLS=$(shell for tool in `find */main.go -type f`; do dirname $$tool; done)
PROGRAMS=$(shell for tool in `find */main.go -type f`; do echo bin/`dirname $$tool`; done)

all: $(PROGRAMS) $(DOCS)
.PHONY: all

bin/%:
	@mkdir -p bin
	@cd $(@F) && go build -o ../bin/$(@F)

install: $(PROGRAMS)
	@for dir in $(INSTALL_DIRS); do install $$dir $(PREFIX)/$$dir; done
.PHONY: install

uninstall:
	@for file in $(INSTALL_FILES); do rm -rf $$file; done
.PHONY: uninstall

share/man/man1:
	@mkdir -p share/man/man1
share/man/man1/%.1: share/man/man1
	@ronn --date="2014-11-01" --manual="User Manual" --organization="$(ORG)" --style=dark man/$(@F).md
	@mv man/$(@F) share/man/man1/$(@F)
	@mv man/$(@F).html docs/$(@F).html

clean:
	@rm -rf bin share dist
.PHONY: clean

dist: $(PROGRAMS)
	@mkdir -p dist
	@tar -czvf dist/cipherutils.tar.gz ./bin ./share ./README.md ./Makefile
.PHONY: dist

distclean:
	@rm -rf dist
.PHONY: distclean

mostlyclean:
	@rm -rf bin share
.PHONY: mostlyclean

check:
	@for tool in $(TOOLS); do cd $$tool && go test; done
.PHONY: check
