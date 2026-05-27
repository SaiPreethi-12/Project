# FROM golang:latest AS builder

# WORKDIR /app

# COPY . .

# RUN go build -o main .

# FROM scratch

# COPY --from=builder /app/main /main

# CMD ["/main"]
# Builder stage
FROM golang:latest AS builder

# Set working directory
WORKDIR /app

# Copy project files
COPY . .

# Build Go application
RUN go build -o main .

# Final lightweight image
FROM scratch

# Copy built binary
COPY --from=builder /app/main /main

# Run application
CMD ["/main"]