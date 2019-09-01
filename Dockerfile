FROM alpine

COPY server /root/

EXPOSE 8080
WORKDIR /root/

CMD ["./server"]
