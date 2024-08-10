# Image Recognition Service

A scalable Image Recognition Service built with Go and Python. This project integrates a pre-trained MobileNetV2 model for high-accuracy object detection and classification. It provides a RESTful API for uploading images and receiving recognition results, with Docker for seamless deployment.

## Features

- **Image Upload**: Handle file uploads via a RESTful API.
- **Image Recognition**: Uses MobileNetV2 model for object detection and classification.
- **Docker Integration**: Containerized application for easy deployment and scalability.
- **High Accuracy**: Achieves over 90% accuracy on standard benchmark datasets.

## Technologies

- **Go**: For the API and server logic.
- **Python**: For running the TensorFlow model.
- **TensorFlow**: For image recognition with MobileNetV2.
- **Docker**: For containerization and deployment.

## Project Structure

```
image-recognition-service/
│
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── handlers/
│   ├── handlers.go
│   ├── image_upload.go
│   ├── image_recognition.go
│
├── model/
│   ├── model.py
│   └── requirements.txt
├── utils/
│   └── utils.go
└── tests/
    ├── handlers_test.go
    └── utils_test.go
```

## Setup and Installation

### Prerequisites

- Go (1.18 or higher)
- Python (3.9 or higher)
- Docker

### Go Application

1. **Initialize Go Modules:**

   ```bash
   go mod tidy
   ```

2. **Build the Go Application:**

   ```bash
   go build -o image-recognition-service
   ```

### Python Environment

1. **Navigate to the `model` Directory:**

   ```bash
   cd model
   ```

2. **Install Python Dependencies:**

   ```bash
   pip install -r requirements.txt
   ```

### Docker

1. **Build the Docker Image:**

   ```bash
   docker build -t image-recognition-service .
   ```

2. **Run the Docker Container:**

   ```bash
   docker run -p 8080:8080 image-recognition-service
   ```

## API Endpoints

### Upload Image

- **Endpoint:** `POST /upload`
- **Form Data:** `image` (file)

**Response:**

```
File uploaded successfully: <filename>
```

### Recognize Image

- **Endpoint:** `POST /recognize`
- **Form Data:** `image` (file)

**Response:**

```
Recognition result: <result>
```

## Testing

To run tests for the Go application:

```bash
go test ./...
```

## Contributing

Feel free to submit issues or pull requests. Contributions are welcome!

## License

This project is licensed under the [MIT License](LICENSE).

---
