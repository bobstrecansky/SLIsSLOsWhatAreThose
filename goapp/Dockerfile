FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go mod vendor
RUN go build -mod vendor -o slislo ./cmd

CMD ["./slislo"]
