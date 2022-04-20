FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o todo-api

EXPOSE 1324

CMD ./todo-api