FROM golang:1.21-alpine AS build

# Setting current working directory
WORKDIR /app

# Copy source code to working directory
COPY . .

# Download dependencies
RUN go mod download && go mod verify

RUN CGO_ENABLED=0 go build -o ./main


# Executing the container in Alpine Linux as base image only -> light weight
FROM alpine:3.14

WORKDIR /app
COPY --from=build /app/ /app/

# Debug /app/main
RUN ls -lrt /app/main
RUN ls -l /app/main

# Exposing port
EXPOSE 9000

# Command to run the executable
CMD ["/app/main"]