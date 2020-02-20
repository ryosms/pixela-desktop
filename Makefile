XGO_OUT_DIR=build/tmp
APP_NAME=pixela-desktop

.PHONY: win darwin linux all clean

all: win mac-gio mac-shiny mac-metal linux

win:
	mkdir -p build/$(APP_NAME)_Windows_x86_64
	xgo --targets=windows/amd64 -tags="nucular-shiny" -x -dest $(XGO_OUT_DIR) -ldflags="-H windowsgui -w -s" ./
	find $(XGO_OUT_DIR) -type f -exec cp {} build/ \;
	cp build/$(APP_NAME)-windows-4.0-amd64.exe build/$(APP_NAME)_Windows_x86_64/$(APP_NAME).exe

mac-gio:
	mkdir -p build/$(APP_NAME)_macOS_x86_64
	xgo --targets=darwin/amd64 -tags="nucular-gio" -dest $(XGO_OUT_DIR) -ldflags="-w -s" ./
	find $(XGO_OUT_DIR) -type f -exec cp {} build/ \;
	cp build/$(APP_NAME)-darwin-*-amd64 build/$(APP_NAME)_macOS_x86_64/$(APP_NAME)_gio

mac-shiny:
	mkdir -p build/$(APP_NAME)_macOS_x86_64
	xgo --targets=darwin/amd64 -tags="nucular-shiny" -dest $(XGO_OUT_DIR) -ldflags="-w -s" ./
	find $(XGO_OUT_DIR) -type f -exec cp {} build/ \;
	cp build/$(APP_NAME)-darwin-*-amd64 build/$(APP_NAME)_macOS_x86_64/$(APP_NAME)_shiny

mac-metal:
	mkdir -p build/$(APP_NAME)_macOS_x86_64
	xgo --targets=darwin/amd64 -tags="metal" -dest $(XGO_OUT_DIR) -ldflags="-w -s" ./
	find $(XGO_OUT_DIR) -type f -exec cp {} build/ \;
	cp build/$(APP_NAME)-darwin-*-amd64 build/$(APP_NAME)_macOS_x86_64/$(APP_NAME)_metal

linux:
	mkdir -p build/$(APP_NAME)_Linux_x86_64
	xgo --targets=linux/amd64 -tags="nucular-shiny" -dest $(XGO_OUT_DIR) -ldflags="-w -s" ./
	find $(XGO_OUT_DIR) -type f -exec cp {} build/ \;
	cp build/$(APP_NAME)-linux-amd64 build/$(APP_NAME)_Linux_x86_64/$(APP_NAME)

clean:
	rm -fr build/*
