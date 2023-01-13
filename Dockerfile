#Build Image
FROM golang:1.19.5 AS builder

#Go to app folder and copy all project files
WORKDIR /app
COPY . .

#Build binary
RUN go build ./cmd/main/main.go

#Run Image
FROM busybox:latest

#Copy Binary from builder to run image
COPY --from=builder /app/main /app/

#Run Binary
CMD [ "/app/main" ]
