FROM golang:alpine
WORKDIR /root/simple_ws_http_server
COPY . .
#RUN apk add --no-cache git
RUN go get -d -v "github.com/gorilla/websocket"
RUN go build -o /root/simple_ws_http_server/serv /root/simple_ws_http_server/serv.go

# FROM alpine:latest
# WORKDIR /root/simple_ws_http_server
# COPY --from=0 /root/simple_ws_http_server/serv .
# COPY --from=0 /root/simple_ws_http_server/static/ .
CMD ["./serv"]