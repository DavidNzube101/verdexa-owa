FROM golang:1.24-alpine AS go-builder
WORKDIR /app
COPY . .
RUN go build -o app

FROM python:3.9-slim
WORKDIR /app
COPY --from=go-builder /app/app .
COPY requirements.txt server.py ./

# Install Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Run the Go binary which starts the Python server
CMD ["./app"]