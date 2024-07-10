run: build
	./build/yerd

build: clean
	go build -o ./build/yerd cmd/main.go

clean:
	rm -f ./build/yerd