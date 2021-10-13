FROM golang:latest as builder
WORKDIR /app
ADD . /app/
RUN go build -o docker-golang

# FROM alpine/latest
FROM scratch
WORKDIR /app
COPY --from=builder /app/docker-golang .

# ARG OS
# ARG ARCH
# RUN env GOOS=$OS goarch=$ARCH go build -o docker-golang .

CMD ["./docker-golang"]