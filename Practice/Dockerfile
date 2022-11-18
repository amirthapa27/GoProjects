FROM golang:1.19.3-alpine3.16

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /practice

EXPOSE 8080

CMD [ "/practice" ]