TARGET := kosmos-reborn-builder
SOURCE := ./src/

all: clean deps $(TARGET)

clean:
	@rm ./$(TARGET)
	@echo Cleaned previous build...

deps:
	@go mod download
	@echo Dependencies downloaded...

$(TARGET):
	@go build -o $@ $(SOURCE)
	@echo Kosmos Reborn Builder has been built - $@
