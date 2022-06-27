build:
	go build -o ./voe-dl src/main.go src/core.go src/utils.go 

install:
	mv ./voe-dl /usr/local/bin