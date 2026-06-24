# syntax=docker/dockerfile:1
#
# Resolver image (HLD §4.4 / dependencies: "single statically-linked binary
# plus a Docker image"). Build context is the resolver/ module directory.

# ---- build stage ----
FROM golang:1.25 AS build
WORKDIR /src
# Cache module downloads on go.mod/go.sum changes only.
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
# Fully static, self-contained binary — the ZK artifacts are go:embed'd, so no
# files need to ship alongside it.
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/resolver ./cmd/resolver

# ---- runtime stage ----
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/resolver /usr/local/bin/resolver
# REST, UDP query port, BitTorrent listen port.
EXPOSE 8080 5353/udp 42069
ENTRYPOINT ["/usr/local/bin/resolver"]
