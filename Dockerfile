FROM golang:1.6-alpine

COPY . /go/src/github.com/rebuy-de/kubernetes-aws-volume-discovery

RUN set -x \
 && cd /go/src/github.com/rebuy-de/kubernetes-aws-volume-discovery \
 && apk add \
        --no-cache \
        ca-certificates \
 && apk add \
        --no-cache \
        --virtual .build-deps \
        git \
 && go get github.com/Masterminds/glide \
 && hack/build.sh \
 && mv kubernetes-aws-volume-discovery /bin/ \
 && apk del .build-deps \
 && rm -rf /go

ENTRYPOINT ["/bin/kubernetes-aws-volume-discovery"]
