#!/usr/bin/env sh

package=cmd/service.go
package_name=otter-identity
version=0.0.1
build_dir=./dist
# For cross-compilation
CGO=0

: "${SERVICE_ENV:=dev}"

output_name=$package_name
if [ "$GOOS" = "windows" ]; then
  output_name+='.exe'
fi

if [ "$SERVICE_ENV" = "dev" ]; then
  ld="-X github.com/otter-im/identity-service/pkg.Version=$(git rev-parse HEAD)"
  tag='-tags dev'
else
  ld="-X github.com/otter-im/identity-service/pkg.Version=v${version}"
  tag=''
fi

env CGO_ENABLED=$CGO GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "${ld}" $tag -trimpath -o ${build_dir}/${output_name} $package

if [ $? -ne 0 ]; then
  exit 1
fi
