FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build ./cmd/main/main.go

EXPOSE 8080

CMD [ "./main" ]
