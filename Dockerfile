FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build your Go app
RUN go build -o foxtrot

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./foxtrot"]

