# Import go image
FROM golang:1.14-alpine
# Label for maintainer
LABEL maintainer="Jack Maarek"
# Set the working directory inside the container
WORKDIR /go/src
# Copy the full project to currennt directory
COPY . .
# Run command to nstall the dependencies
RUN go install

CMD go run main.go

EXPOSE 8080