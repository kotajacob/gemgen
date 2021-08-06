# gemgen
# See LICENSE for copyright and license details.
.POSIX:

include config.mk

all: clean build

build:
	go build -ldflags "-X main.Version=$(VERSION)"
	scdoc < gemgen.1.scd | sed "s/VERSION/$(VERSION)/g" > gemgen.1

clean:
	rm -f gemgen
	rm -f gemgen.1

install: build
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f gemgen $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/gemgen
	mkdir -p $(DESTDIR)$(MANPREFIX)/man1
	cp -f gemgen.1 $(DESTDIR)$(MANPREFIX)/man1/gemgen.1
	chmod 644 $(DESTDIR)$(MANPREFIX)/man1/gemgen.1

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/gemgen
	rm -f $(DESTDIR)$(MANPREFIX)/man1/gemgen.1

.PHONY: all build clean install uninstall
