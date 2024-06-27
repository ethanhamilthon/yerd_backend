run: build
	./build/word

build: clean
	go build -o ./build/word cmd/main.go

clean:
	rm -f ./build/word