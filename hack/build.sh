#/bin/bash

set -ex

VERSION=$( git describe --always --dirty | tr '-' '.' )

glide install

go test \
	-p 1 \
	$(glide novendor)

go build \
	-ldflags "-X main.version=${VERSION}"
