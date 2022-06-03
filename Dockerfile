FROM golang:1.18.3-alpine3.16

WORKDIR /go/src/hex-arch-go

COPY . .

RUN apk add git
RUN go get
RUN go build -o /go/bin/hex-arch-go .

CMD [ "/go/bin/hex-arch-go" ]