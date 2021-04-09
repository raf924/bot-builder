FROM golang:1.16.3-alpine as builder
#RUN apk add --no-cache git
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
WORKDIR /go/src/github.com/raf924/bot-builder
COPY go.mod go.sum /go/src/github.com/raf924/bot-builder/
RUN go mod download
COPY . /go/src/github.com/raf924/bot-builder/
RUN CGO_ENABLED=0 go build -o /bot-builder ./cmd/bot-builder
FROM scratch
COPY --from=builder /bot-builder /
ENTRYPOINT ["/bot-builder"]