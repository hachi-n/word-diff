APP_NAME := word-diff

go-clean:
	rm -f ./build/$(APP_NAME)

go-build:
	go build -o ./build/$(APP_NAME) ./cli/$(APP_NAME)

copy-tool:
	cp -pr  ./build/$(APP_NAME) /usr/local/bin/$(APP_NAME)

remove-tool:
	rm -f /usr/local/bin/$(APP_NAME)

build:
	$(MAKE) go-clean
	$(MAKE) go-build

clean:
	$(MAKE) go-clean

install:
	$(MAKE) go-clean
	$(MAKE) go-build
	$(MAKE) copy-tool

uninstall:
	$(MAKE) remove-tool

