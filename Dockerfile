FROM golang AS builder
WORKDIR /go/src/go-libvirt-web
COPY . .
RUN apt-get update && apt-get install -y libvirt-dev
RUN go get -d .
RUN go build .

FROM scratch
WORKDIR /app
COPY --from=builder /go/src/go-libvirt-web/libvirt-web .
ENTRYPOINT [ "/app/libvirt-web" ]