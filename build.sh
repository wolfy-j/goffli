#!/bin/bash
VERSION=1.0

cd $(dirname "${BASH_SOURCE[0]}")
OD="$(pwd)"

package(){
	echo Packaging $1 Binary

	bdir=goffli-${VERSION}-$2-$3
	rm -rf package/$bdir && mkdir -p package/$bdir
	GOOS=$2 GOARCH=$3 ./build.sh

	if [ "$2" == "windows" ]; then
		mv goffli package/$bdir/goffli.exe
	else
		mv goffli package/$bdir
	fi

	cp README.md package/$bdir
	cp CHANGELOG.md package/$bdir
	cp LICENSE package/$bdir
	cd package

	if [ "$2" == "linux" ]; then
		tar -zcf $bdir.tar.gz $bdir
	else
		zip -r -q $bdir.zip $bdir
	fi

	rm -rf $bdir
	cd ..
}

if [ "$1" == "package" ]; then
	rm -rf package/
	package "Windows" "windows" "amd64"
	package "Mac" "darwin" "amd64"
	package "Linux" "linux" "amd64"
	package "FreeBSD" "freebsd" "amd64"
	exit
fi

CGO_ENABLED=0 go build -ldflags "$LDFLAGS -extldflags '-static'" -o "$OD/goffli" main.go