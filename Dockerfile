FROM golang:1.20

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o binary

EXPOSE 8080

ENTRYPOINT ["/app/binary"]