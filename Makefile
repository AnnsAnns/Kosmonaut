TARGET := builder
SOURCE := ./cmd/builder/
OUTPUT := ./build/

all: clean deps $(TARGET)

clean:
	@rm -rf ./$(OUTPUT)
	@echo Cleaned previous build...

deps:
	@go mod download
	@echo Dependencies downloaded...

$(TARGET):
	@mkdir -p ./$(OUTPUT)/modules
	@cp -R ./assets/* ./$(OUTPUT)/
	@go build -o ./$(OUTPUT)/$@ $(SOURCE)
	@echo Kosmos Reborn Builder has been built - $(OUTPUT)$@
