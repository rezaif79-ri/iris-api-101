FROM golang:1.21-alpine

RUN apk add --no-cache git

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/app ./app/main.go

EXPOSE 3000

CMD ["./out/app"]