FROM golang:latest as build

RUN mkdir -p /go/src/greet
WORKDIR /go/src/greet

RUN go get -d github.com/gorilla/mux && \
	go get -d github.com/prometheus/client_golang/prometheus && \
    go get -d github.com/prometheus/client_golang/prometheus/promhttp

COPY main.go . 

RUN CGO_ENABLED=0 go build -a -installsuffix cgo --ldflags "-s -w" -o /usr/bin/server 

FROM alpine

COPY --from=build /usr/bin/server /root/

EXPOSE 8080
WORKDIR /root/

CMD ["./server"]
