
INSTALL_DEST=/usr/local/bin

setup:
	go get .

run:
	go run .

build:
	go build

install:
	@echo "Installing to ${INSTALL_DEST}"
	@sudo mv -f ./openpomodoro-cli ${INSTALL_DEST}
