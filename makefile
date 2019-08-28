# Variables
BINARY_NAME=main

# Commands
clean: 
	rm -f $(BINARY_NAME) main.zip

build: 
	GOOS=linux go build -o $(BINARY_NAME)

zip: 
	zip -r main.zip $(BINARY_NAME)

deploy: clean build zip