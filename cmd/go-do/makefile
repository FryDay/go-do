GOCMD=go
TAR=tar zcf
ZIP=zip -j
BIN=bin
OUT=go-do

LINUX=GOOS=linux GOARCH=amd64
WINDOWS=GOOS=windows GOARCH=amd64
MAC=GOOS=darwin GOARCH=amd64

all: linux mac windows

linux:
	$(LINUX) $(GOCMD) build -o $(BIN)/linux/$(OUT) $(MAIN)
	$(TAR) $(BIN)/$(OUT)-linux.tar.gz --directory=$(BIN)/linux $(OUT)

mac:
	$(MAC) $(GOCMD) build -o $(BIN)/mac/$(OUT) $(MAIN)
	$(TAR) $(BIN)/$(OUT)-mac.tar.gz --directory=$(BIN)/mac $(OUT)

windows:
	$(WINDOWS) $(GOCMD) build -o $(BIN)/windows/$(OUT).exe $(MAIN)
	$(ZIP) $(BIN)/$(OUT)-windows.zip $(BIN)/windows/$(OUT).exe
