
#include $(GOROOT)/src/Make.inc

all:	install

install:
	cd raster && make install
	cd curve && make install

clean:
	cd raster && make clean
	cd curve && make clean

nuke:
	cd raster && make nuke
	cd curve && make nuke

fmt:
	gofmt -w . 

