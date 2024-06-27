FROM golang:1.22-alpine AS builder-backend

WORKDIR /build

COPY go.mod go.sum ./  
RUN go mod download 

COPY ./ ./                   
RUN go build -o /build/yerd cmd/main.go

FROM gcr.io/distroless/base-debian12 AS backend

WORKDIR /app
COPY --from=builder-backend /build/yerd .
COPY .env /app/.env

CMD ["./yerd"]