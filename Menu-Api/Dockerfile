FROM golang:1.19.3-alpine3.16

WORKDIR /app

COPY go.mod /.
COPY go.sum /.

RUN go mod download

COPY . .

RUN go build -o /menu-api

COPY . .

EXPOSE 8080

CMD [ "/menu-api" ]