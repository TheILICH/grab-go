FROM golang:1.22.2-alpine3.19 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

FROM alpine:3.19

WORKDIR /app

COPY wait.sh /app/
RUN chmod +x /app/wait.sh

RUN apk --no-cache add postgresql-client

# Copy the compiled binary
COPY --from=builder /app/myapp /app/

# Copy templates and static files into the final image
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/views /app/views
# If you have an assets directory, include it as well
COPY --from=builder /app/assets /app/assets
COPY --from=builder /app/.env /app/.env

EXPOSE 8080

# CMD ["/app/myapp"]
