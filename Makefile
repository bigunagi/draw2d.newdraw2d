
#include $(GOROOT)/src/Make.inc

all:	install

build:
	cd raster && go build
	cd curve && go build
	cd geometry && go build
	
install:
	cd raster && go install
	cd curve && go install
	cd geometry && go install
	
clean:
	cd raster && go clean
	cd curve && go clean
	cd geometry && go clean

fmt:
	gofmt -w . 

