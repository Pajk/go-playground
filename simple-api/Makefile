
start: build
	./bin/simple-api

build: clean
	go build -o bin/simple-api

watch:
	go get github.com/codegangsta/gin
	gin run app.go

clean:
	rm -rf ./bin

docker-build:
	docker build -t go-playground/simple-api .

docker-start:
	docker-compose up --build

.PHONY: build watch
