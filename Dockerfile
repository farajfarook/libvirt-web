FROM golang AS builder
WORKDIR /go/src/libvirt-web
COPY . .
RUN apt-get update && apt-get install -y libvirt-dev
RUN go get -d . && go build .

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/libvirt-web/libvirt-web .
ENTRYPOINT [ "/app/libvirt-web" ]