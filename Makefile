all:
	@./build.sh
clean:
	rm -f goffli
install: all
	cp goffli /usr/local/bin/goffli
uninstall:
	rm -f /usr/local/bin/goffli
package:
	@./build.sh package
