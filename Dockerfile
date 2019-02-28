FROM golang:1.10

ENV GOPATH=/go
ENV SRV_HOST=0.0.0.0
ENV HTTP_PORT=8080
ENV HTTP_PATH=ht
ENV WS_PORT=8081
ENV WS_PATH=ws
WORKDIR /go/src/cell_radio
COPY ./src/cell_radio/ .

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -d -v ./...
RUN dep ensure
RUN go install -v ./...

CMD cell_radio -host $SRV_HOST -httpPort $HTTP_PORT -httpPath $HTTP_PATH -wsPort $WS_PORT -wsPath $WS_PATH