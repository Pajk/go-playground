FROM golang:1.8

WORKDIR /go/src/github.com/pajk/go-playground/simple-api
RUN go get github.com/codegangsta/gin

COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD gin -p 8000 run app.go