build:
	mkdir -p functions
	go mod download
	go build -o functions/shorten main.go 
