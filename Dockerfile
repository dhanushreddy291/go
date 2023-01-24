FROM golang:buster
WORKDIR /server
COPY main.go main.go
EXPOSE 80
CMD ["go", "run", "main.go"]