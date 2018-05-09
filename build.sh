#!/bin/bash
VERSION=1.0

build(){
	echo Packaging $1 Binary

	bdir=goffli-${VERSION}-$2-$3
	rm -rf build/$bdir && mkdir -p build/$bdir
	GOOS=$2 GOARCH=$3 ./build.sh

	if [ "$2" == "windows" ]; then
		mv goffli build/$bdir/goffli.exe
	else
		mv goffli build/$bdir
	fi

	cp README.md build/$bdir
	cp CHANGELOG.md build/$bdir
	cp LICENSE build/$bdir
	cd build

	if [ "$2" == "linux" ]; then
		tar -zcf $bdir.tar.gz $bdir
	else
		zip -r -q $bdir.zip $bdir
	fi

	rm -rf $bdir
	cd ..
}

if [ "$1" == "package" ]; then
	rm -rf build/
	build "Windows" "windows" "amd64"
	build "Mac" "darwin" "amd64"
	build "Linux" "linux" "amd64"
	build "FreeBSD" "freebsd" "amd64"
	exit
fi