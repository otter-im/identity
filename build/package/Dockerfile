FROM golang:1.17-alpine

ARG SERVICE_ENV=dev
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH

RUN apk add git

COPY . ./build
WORKDIR ./build

RUN go mod download && \
    go mod verify

RUN scripts/build_container.sh

FROM alpine:latest
COPY --from=0 /go/build/dist/identity-service /usr/bin/identity-service
CMD ["/usr/bin/identity-service", "-logtostderr"]
