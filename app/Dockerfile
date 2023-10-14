# 開発環境用
FROM golang:1.20 AS develop

WORKDIR /go/src/github.com/mahiro72/tomocafe-api

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml.develop"]

## Build
FROM golang:1.20 AS build

WORKDIR /go/src/github.com/copalette/copalette-api

COPY app .

RUN go mod download

COPY db db

COPY .env .env

RUN [ ! -f .env ] || export $(grep -v '^#' .env | xargs)

RUN CGO_ENABLED=0 go build -o /server .

## Deploy
FROM golang:1.20 AS deploy

WORKDIR /

COPY --from=build /server /server

EXPOSE 8080

CMD ["/server"]
