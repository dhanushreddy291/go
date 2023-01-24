FROM golang:latest
WORKDIR /server
COPY main.go main.go
EXPOSE 80
CMD ["go", "run", "main.go"]