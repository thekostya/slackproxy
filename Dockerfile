FROM golang
WORKDIR /go/src/github.com/thekostya/test1
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/thekostya/test1/main .
CMD ["./main"]
