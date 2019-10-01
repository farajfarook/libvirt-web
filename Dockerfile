FROM golang:alpine AS builder
WORKDIR /go/src/libvirt-web
COPY . .
RUN apk add --update libvirt-dev git gcc g++
RUN go get github.com/swaggo/swag/cmd/swag && swag init
RUN go get -d . && go build .

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/libvirt-web/libvirt-web . 
RUN apk add --update libvirt
ENTRYPOINT [ "/app/libvirt-web" ]