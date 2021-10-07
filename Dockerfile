FROM registry.redhat.io/rhel8/go-toolset:latest@sha256:db74be244cbf62081667253dbafeb0af75c3410982a45767d1dca6a3eb86f49b AS builder
WORKDIR /opt/app-root/src
COPY . .
RUN go build -o bin/festoji main.go

FROM scratch
COPY --from=builder /opt/app-root/src/bin/festoji /usr/bin/festoji
CMD ["festoji"]
