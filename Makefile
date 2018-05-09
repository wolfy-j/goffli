all:
	@./build.sh
clean:
	rm -f goffli
install: all
	cp goffli /usr/local/goffli
uninstall:
	rm -f /usr/local/bin/goffli
package:
	@./build.sh package