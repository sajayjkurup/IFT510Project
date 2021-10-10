FROM --platform=$BUILDPLATFORM golang:alpine AS builder
ARG TARGETARCH
ARG BUILDPLATFORM   
ARG TARGETPLATFORM
ARG TARGETOS
RUN echo "Building on $BUILDPLATFORM, building for $TARGETPLATFORM"
WORKDIR /app
COPY . /app/.
RUN env GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/target/caseconv

FROM scratch
LABEL org.opencontainers.image.source="https://github.com/sajayjkurup/IFT510Project"
COPY --from=builder /app/target/caseconv /main
ENTRYPOINT ["/main"]