# Build stage
FROM golang:latest as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend/apiserver/out/apiserver backend/apiserver/cmd/apiserver/main.go

# Runtime stage
FROM golang:latest

WORKDIR /app

COPY --from=build /app/backend/apiserver/out/apiserver /app/apiserver

EXPOSE 8080

CMD ["./apiserver"]
