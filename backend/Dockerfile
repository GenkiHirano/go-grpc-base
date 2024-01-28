FROM golang:1.21 AS builder

ARG TARGETOS=linux
ARG TARGETARCH=amd64

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY internal/ internal/

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -o server ./internal/main.go

FROM gcr.io/distroless/static:nonroot

ENV TZ=Asia/Tokyo
LABEL org.opencontainers.image.title="ai-fortune/server" \
    org.opencontainers.image.description="ai-fortune server" \
    org.opencontainers.image.url="https://github.com/blueish-co/ai-fortune" \
    org.opencontainers.image.vender="Blueish Co., Ltd."

WORKDIR /
COPY --from=builder /app/server /server
USER 65532:65532

ENTRYPOINT ["/server"]
