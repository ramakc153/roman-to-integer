FROM golang:1.21

WORKDIR /app

COPY go.mod

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /roman-to-integer

EXPOSE 8080
CMD ["/roman-to-integer"]

