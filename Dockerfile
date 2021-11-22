FROM golang:alpine
#port
COPY . .
EXPOSE 8080
# run api
RUN go run main.go
