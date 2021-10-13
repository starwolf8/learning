FROM golang:latest as builder
WORKDIR /app
ADD . /app/
# Include command CGO_ENABLED=0 to build statically linked images
RUN CGO_ENABLED=0 go build -o docker-golang

# FROM alpine/latest
FROM scratch
WORKDIR /app
COPY --from=builder /app/docker-golang .

# ARG OS
# ARG ARCH
# RUN env GOOS=$OS goarch=$ARCH go build -o docker-golang .

CMD ["./docker-golang"]