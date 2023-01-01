# syntax=docker/dockerfile=1

FROM golang:1.19.4-alpine
WORKDIR /simple-crud
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./cmd ./cmd
RUN echo "simple-crud service started"
EXPOSE 8080
CMD ["/simple-crud/cmd"]