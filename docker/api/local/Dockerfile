FROM golang:1.24-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.47.0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air"]
