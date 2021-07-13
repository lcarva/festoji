FROM registry.access.redhat.com/ubi8/ubi:latest
COPY ./festoji /usr/bin/
CMD ["festoji"]
