FROM golang:1.18-rc-alpine3.14
WORKDIR /app
COPY . .
RUN go mod download
RUN apk update && \
	apk add net-tools
RUN go build -o main
ENTRYPOINT ["./main"]