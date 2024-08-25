all: build run

build:
	go mod init greport.com/m/v2 || true
	go mod tidy
	go build -o scanner
	
run:
	./scanner

clean:
	rm -f scanner

fclean: clean
	rm -rf ./logs