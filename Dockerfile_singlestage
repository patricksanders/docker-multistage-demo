FROM golang:1.11

WORKDIR /go/src/app/

COPY *.go ./
ENTRYPOINT [ "/go/src/app/hello" ]

RUN go test -v -cover
RUN go build -ldflags "-extldflags \"-static\"" -o hello
