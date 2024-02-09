# Go HTTP Server with Worker

This project demonstrates a simple Go HTTP server that receives JSON requests and processes them asynchronously using a worker goroutine. Each request is converted to a different format and sent to a webhook.

## How it works

The Go HTTP server listens for incoming requests on port 8080. When a request is received, it decodes the JSON body into a struct, processes it asynchronously, and responds with an acknowledgment message.

The worker goroutine converts each incoming request into a different format, following the provided specifications. It then sends the converted request to a predefined webhook URL.

## Instructions to run

1. **Install Go**: Make sure you have Go installed on your system. If not, download and install it from the [official website](https://golang.org/doc/install).

2. **Clone the repository**: Clone this repository to your local machine:

   ```bash
   git clone https://github.com/YashTripathi01/customerlabs-interview.git
   ```

3. **Navigate to project directory**: Change directory to the project directory:

   ```bash
   cd customerlabs-interview
   ```

4. **Run the server**: Execute the following command to run the Go program:

   ```bash
   go run main.go
   ```

5. **Access the server**: Use an HTTP client like [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/) to send JSON requests to the server [endpoint](http://localhost:8080/) and trigger the processing

6. **Send requests**: Send JSON requests to the server endpoint (e.g., using cURL or Postman) to trigger the processing.

7. **Check output**: The processed requests will be printed in the terminal where you ran the Go program.

## Dependencies

This project uses only standard Go packages and does not require any additional dependencies.

## Author

Yash Tripathi
