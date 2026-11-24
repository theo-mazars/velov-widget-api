# syntax=docker/dockerfile:1

# ---- Build stage ----
# golang:1.26.11-alpine
FROM --platform=$BUILDPLATFORM golang:1.26.11-alpine@sha256:ecb67f987aa1a57794c91d19ade7ae0209f2f0f5496aaef9e6ae85ddf92d02ab AS builder

WORKDIR /src

# Download dependencies first to benefit from layer caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY api ./api

# Build a static binary for the target platform (amd64 / arm64)
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags="-s -w" -o /out/velov-widget-api ./api

# ---- Runtime stage ----
# Distroless static: no shell, no package manager, CA certs + tzdata included, non-root user
# gcr.io/distroless/static-debian12:nonroot
FROM gcr.io/distroless/static-debian12:nonroot@sha256:afa5c872c891853ca7fcf1f12c3edb23f7eeef36189728842dd51042ff57f7ab

WORKDIR /app

COPY --from=builder --chown=root:root --chmod=0555 /out/velov-widget-api ./velov-widget-api
COPY --chown=root:root --chmod=0444 resources ./resources

USER nonroot:nonroot

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
    CMD ["/app/velov-widget-api", "-healthcheck"]

ENTRYPOINT ["/app/velov-widget-api"]
