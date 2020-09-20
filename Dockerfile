FROM golang:1.15.0-alpine

WORKDIR /usr/src/app
COPY . .
RUN go build -o main .
CMD ["./main"]