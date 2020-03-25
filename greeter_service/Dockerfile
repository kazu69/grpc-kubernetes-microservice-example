FROM golang

COPY . /go/src/greeter_service
WORKDIR /go/src/greeter_service
RUN go get .
ENTRYPOINT go run main.go
EXPOSE 3000
