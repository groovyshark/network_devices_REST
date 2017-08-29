FROM golang

ADD . /home/rzhevskiy/GoWorkspace/src/networkDevices_REST

RUN go install networkDevices_REST

ENTRYPOINT /go/bin/network_devices_REST

EXPOSE 8080
