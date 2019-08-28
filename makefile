# Variables

# Commands

clean: 
	rm -f main main.zip

build: 
	GOOS=linux go build -o main

zip: 
	zip -r main.zip main

deploy: 
	clean build zip