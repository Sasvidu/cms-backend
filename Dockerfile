# Builder stage
FROM golang:alpine AS builder

WORKDIR /app 

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main

# Final stage
FROM alpine

WORKDIR /app

COPY --from=builder /app/main .

# Copy the .env file
COPY .env .

# Install bash and prepare entrypoint
RUN apk add --no-cache bash && \
    echo "source .env" >> /app/entrypoint.sh && \
    chmod +x /app/entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh", "./main"]
