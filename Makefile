all: build
build: libstorage-mock.so

libstorage-mock.so: *.go mod/*.go
	go build -buildmode plugin -o $@ ./mod

clean:
	rm -f libstorage-mock.so
