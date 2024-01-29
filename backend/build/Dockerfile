FROM golang:1.21 AS builder

ARG TARGETOS=linux
ARG TARGETARCH=amd64

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY internal/ internal/

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -o backend ./internal/cmd/server/main.go

FROM gcr.io/distroless/static:nonroot

ENV TZ=Asia/Tokyo
LABEL org.opencontainers.image.title="go-grpc-base/backend" \
    org.opencontainers.image.description="go-grpc-base backend" \
    org.opencontainers.image.url="https://github.com/GenkiHirano/go-grpc-base" \
    org.opencontainers.image.vender="GenkiHirano"

WORKDIR /
COPY --from=builder /app/backend /backend
USER 65532:65532

ENTRYPOINT ["/backend"]
