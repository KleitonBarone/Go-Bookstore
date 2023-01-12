#Build Image
FROM golang:1.16-alpine AS builder

#Go to app folder and copy all project files
WORKDIR /app
COPY . .

#Build binary
RUN go build ./cmd/main/main.go

#Expose Host Port
EXPOSE 8080

#Run Binary
CMD [ "./main" ]
