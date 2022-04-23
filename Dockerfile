FROM golang:1.17.9

WORKDIR /app

COPY . .

RUN go build -o todo-api

EXPOSE 8080

CMD ./todo-api