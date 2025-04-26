# Build stage
FROM golang:1.21 AS builder

WORKDIR /app
COPY . .
RUN go build -o main .

# Final stage
FROM python:3.9-slim

WORKDIR /app

# Install Python dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy Go binary and Python script
COPY --from=builder /app/main /app/main
COPY your_script.py .

# Run the Go binary (which executes the Python script)
CMD ["./main"]