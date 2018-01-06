.PHONY: build clean

default: build

OUTPUT_PATH=build/goReminder

clean:
	rm -rf build && mkdir build && touch build/.gitkeep

build: clean
	go build -o $(OUTPUT_PATH)

run:
	go run index.go

goRun:
	$(OUTPUT_PATH)