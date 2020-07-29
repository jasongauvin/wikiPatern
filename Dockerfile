# Import go image
FROM golang:1.14-alpine
# Label for maintainer
LABEL maintainer="Jack Maarek"
# Set the working directory inside the container
WORKDIR /go/src
# Copy the full project to currennt directory
COPY . .
# Set env to fix error with go tests
ENV CGO_ENABLED 0
# Run command to install the dependencies
RUN go mod download

EXPOSE 8080