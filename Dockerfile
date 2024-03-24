FROM golang:1.22.1-alpine3.19 as builder
ENV CGO_ENABLED=0
WORKDIR /tmp/buildspace
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go mod tidy
RUN go build -o /main .

# a new clean image with just the binary
FROM alpine:3.19
COPY --from=builder /main /main
CMD ["/main"]
