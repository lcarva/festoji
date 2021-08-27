FROM registry.redhat.io/rhel8/go-toolset:latest AS builder
WORKDIR /opt/app-root/src
COPY . .
RUN go build -o bin/festoji main.go

FROM scratch
COPY --from=builder /opt/app-root/src/bin/festoji /usr/bin/festoji
CMD ["festoji"]
