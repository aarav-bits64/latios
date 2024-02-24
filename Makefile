GC = go
SRC_DIR = src
LATIOS_GO = $(SRC_DIR)/latios.go 
OUTPUT_BIN = latios

build:
	$(GC) build -o $(OUTPUT_BIN) $(LATIOS_GO)

install:
	sudo mv $(OUTPUT_BIN) /usr/local/bin/

.PHONY: build