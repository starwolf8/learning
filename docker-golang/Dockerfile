FROM golang:latest as builder
WORKDIR /app
ADD . /app/
# Include command CGO_ENABLED=0 to build statically linked images
RUN CGO_ENABLED=0 go build -o docker-golang

# FROM alpine/latest
FROM scratch
WORKDIR /app

COPY --from=builder /app/docker-golang .
# COPY --from=builder /app/static/* ./static/
# COPY --from=builder /app/templates/* ./templates/
# COPY --from=builder /app/app.env .

# ARG OS
# ARG ARCH
# RUN env GOOS=$OS goarch=$ARCH go build -o docker-golang .

CMD ["./docker-golang"]