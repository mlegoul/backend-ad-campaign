FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o backend-ad-campaign cmd/rest/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/backend-ad-campaign /backend-ad-campaign

EXPOSE 8080

CMD ["/backend-ad-campaign"]
