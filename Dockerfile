FROM golang:1.19.8 as builder
ENV GOPROXY=https://goproxy.cn
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ARG TARGETOS
ARG TARGETARCH

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the go source
COPY cmd/ cmd/
COPY pkg/ pkg/

ARG VERSION

# Compile all the binaries
RUN go build -v -o hyperos_backend_demo cmd/main.go

#RUN upx hyperos_backend_demo 
#
# IMAGE TARGETS
# -------------
FROM gengweifeng/gcr-io-distroless-static-nonroot as hyperos_backend_demo
WORKDIR /
COPY --from=builder /workspace/hyperos_backend_demo .
USER 65532:65532
EXPOSE 8001
ENTRYPOINT ["/hyperos_backend_demo"]
