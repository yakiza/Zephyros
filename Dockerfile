FROM golang:1.16-alpine

WORKDIR /opt/zephyros

RUN export GO111MODULE=on
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ /opt/zephyros/

RUN cd cmd  && go build -o zephyros

EXPOSE 8081

CMD ["/opt/zephyros/cmd/zephyros"]
