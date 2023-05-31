FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

ENV PATH="/root/go/bin:${PATH}"

COPY . .

EXPOSE 3001

CMD ["air"]
