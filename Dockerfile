FROM golang:1.16.2-alpine as builder
RUN apk add --no-cache ca-certificates git
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
WORKDIR /go/src/github.com/raf924/bot-builder
ENV GO111MODULE=on
COPY go.mod go.sum /go/src/github.com/raf924/bot-builder/
RUN go mod download
COPY . /go/src/github.com/raf924/bot-builder/
RUN CGO_ENABLED=0 go build -o /bot-builder ./cmd/bot-builder
RUN chmod +x /bot-builder
ARG PART
ARG RECIPE
COPY $RECIPE /config/builder-recipe.yaml
RUN mkdir /dist ; /bot-builder ${PART} --recipe "/config/builder-recipe.yaml" -o /dist/main.go
WORKDIR /dist
RUN ls -lsa
RUN go env -w GOPRIVATE=github.com/raf924/*
RUN go mod init localhost/relay
RUN go generate /dist
RUN CGO_ENABLED=0 go build -o /bot /dist/main.go
FROM scratch
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /bot /bot
USER nobody:nobody
VOLUME [ "/config", "/certs", "/data" ]
ENTRYPOINT [ "/bot" ]