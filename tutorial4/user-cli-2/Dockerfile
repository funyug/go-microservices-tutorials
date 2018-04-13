FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/funyug/go-micro-services-tutorials/tutorial4/user-cli-2

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/funyug/go-micro-services-tutorials/tutorial4/user-cli-2 .

ENTRYPOINT ["./user-cli-2"]

CMD ["./user-cli-2"]