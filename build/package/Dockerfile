FROM golang:1.17-alpine

ARG SERVICE_ENV=dev
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH

RUN apk add git

COPY . ./build
WORKDIR ./build

RUN go mod download && \
    go mod verify

RUN scripts/build.sh

FROM alpine:latest
COPY --from=0 /go/build/dist/otter-identity /usr/bin/otter-identity
CMD ["/usr/bin/otter-identity"]
