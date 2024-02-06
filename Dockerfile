FROM golang:1.21 as builder

ARG VERSION
ARG GIT_COMMIT

WORKDIR /app/

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.version=${VERSION} -X main.gitCommit=${GIT_COMMIT}" -o lnd-exporter .

FROM alpine:latest

RUN apk update && apk add ca-certificates

COPY --from=builder /app/lnd-exporter /app/

ENTRYPOINT [ "/app/lnd-exporter" ]
