FROM golang

WORKDIR /app

EXPOSE 8080

ENV SRC_DIR=/go/src/github.com/groovyshark/network_devices_REST

ADD . $SRC_DIR

RUN go get github.com/gorilla/mux
RUN cd $SRC_DIR; go build -o network_rest; cp network_rest /app/

ENTRYPOINT ["./network_rest"]
