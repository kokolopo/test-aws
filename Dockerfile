FROM golang:1.18-alpine3.14

WORKDIR /app

COPY . .

RUN go build -o todo-api

EXPOSE 1324

CMD ./todo-api