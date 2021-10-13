FROM golang:latest

ENV PORT 8080

WORKDIR /app

ADD . /app/

RUN go mod tidy
RUN go mod verify

EXPOSE $PORT

RUN go build -o docker-golang

# ARG OS
# ARG ARCH
# RUN env GOOS=$OS goarch=$ARCH go build -o docker-golang .

CMD ["./docker-golang"]