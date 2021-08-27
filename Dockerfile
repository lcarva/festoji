FROM registry.fedoraproject.org/fedora:36 AS builder

RUN dnf -y install \
    --setopt=deltarpm=0 \
    --setopt=install_weak_deps=false \
    --setopt=tsflags=nodocs \
    golang

WORKDIR /go/src/app
COPY . .
RUN go build -o bin/festoji main.go

FROM scratch
COPY --from=builder /go/src/app/bin/festoji /usr/bin/festoji
CMD ["festoji"]
