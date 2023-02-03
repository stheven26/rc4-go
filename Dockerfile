FROM golang:1.18-alpine
WORKDIR /app
COPY . .
RUN go build -o rc4-go
EXPOSE 8080
CMD ./rc4-go