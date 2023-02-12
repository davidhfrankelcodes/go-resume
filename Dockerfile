# Use the latest version of Go as the base image
FROM golang:latest

# Copy the main.go file into the container
COPY main.go /go/src/app/main.go

# Copy the static files into the container
COPY static /go/src/app/static

# Copy the resume.html file into the container
COPY resume.html /go/src/app/resume.html

# Copy the resume.yaml file into the container
COPY resume.yaml /go/src/app/resume.yaml

# Copy the go.mod file into the container
COPY go.mod /go/src/app/go.mod

# Set the working directory to the app's directory
WORKDIR /go/src/app

# Download the dependencies specified in the go.mod file
# RUN go mod download
RUN go mod download gopkg.in/yaml.v2

# Build the Go app
RUN go build main.go

# Set the command to run when the container starts
CMD ["./main"]

# Expose port 8080 so that the app can be accessed from outside the container
EXPOSE 8080
