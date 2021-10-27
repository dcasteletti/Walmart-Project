FROM golang:1.17 AS BUILDER
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM golang:1.17
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD app/main