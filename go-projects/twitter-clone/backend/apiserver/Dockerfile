# Build stage
FROM golang:latest as build

WORKDIR /app/

COPY ../.. .

RUN go build -o out/apiserver cmd/apiserver/main.go

# Runtime stage
FROM golang:latest

WORKDIR /app/

COPY --from=build /app/backend/apiserver/out/apiserver /app/apiserver

EXPOSE 8080

CMD ["./apiserver"]
