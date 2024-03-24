FROM golang:1.22.1-alpine3.19 as builder
# Setup git to allow go to fetch deps from riptano
#ARG GITHUB_TOKEN
#RUN apk add --no-cache ca-certificates git
#RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
# copy over the code and build it
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
