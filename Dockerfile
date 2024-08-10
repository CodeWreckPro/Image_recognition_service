# Dockerfile
FROM golang:1.18-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /image-recognition-service

# Final image
FROM python:3.9-slim

WORKDIR /app

COPY --from=builder /image-recognition-service /image-recognition-service
COPY model/requirements.txt /app/model/requirements.txt

RUN pip install --no-cache-dir -r /app/model/requirements.txt

COPY model/model.py /app/model/model.py

EXPOSE 8080

CMD ["/image-recognition-service"]
