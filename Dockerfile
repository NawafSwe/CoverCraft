FROM golang:latest

# Set the working directory inside the container
WORKDIR /src

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o main cmd/main.go

# Expose on Port 80
EXPOSE 8080

CMD ["./main"]