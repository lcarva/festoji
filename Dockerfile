FROM registry.redhat.io/rhel8/go-toolset:latest AS builder
COPY . .
RUN go build -o /opt/app-root/src/bin/festoji main.go
FROM scratch
COPY --from=builder /opt/app-root/src/bin/festoji /usr/bin/festoji
CMD ["festoji"]
