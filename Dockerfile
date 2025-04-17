FROM golang:1.24.2

# Install dependencies
RUN apt-get update && apt-get install -y git

# Mark /app as a safe Git directory
RUN git config --global --add safe.directory /app

# Install Air
RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# Use Air to run the app with live reload
CMD ["air"]