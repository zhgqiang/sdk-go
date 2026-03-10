FROM alpine:latest

ADD go-driver-mqtt-demo /app/
ADD etc/config.docker.yaml /app/etc/config.yaml
WORKDIR /app
ENTRYPOINT ["./go-driver-mqtt-demo"]