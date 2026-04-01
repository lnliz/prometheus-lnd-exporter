FROM golang:1.25-alpine AS builder

ARG VERSION
ARG GIT_COMMIT

RUN apk add --no-cache gcc musl-dev

WORKDIR /app/

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags "-linkmode external -extldflags '-static' -X main.version=${VERSION} -X main.gitCommit=${GIT_COMMIT}" -o lnd-exporter .

FROM alpine:latest

RUN apk update && apk add ca-certificates

COPY --from=builder /app/lnd-exporter /app/

ENTRYPOINT [ "/app/lnd-exporter" ]
