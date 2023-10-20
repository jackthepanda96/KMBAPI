FROM golang:1.21-alpine

COPY . /app

WORKDIR /app

RUN echo ${{ secrets.GOOCREDS }} >> credentials.json

RUN go mod tidy

RUN go build -o app .

CMD ["/app/app"]
