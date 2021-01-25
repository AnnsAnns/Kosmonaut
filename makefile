all:
	go mod download
	go build -o builder ./src/
