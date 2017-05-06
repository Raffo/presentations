FROM golang 

ENV GOPATH /go
ENV GOBIN /go/bin
ENV PATH=$GOBIN:$PATH
RUN mkdir -p /go/src/app
ADD main.go /go/src/app/foo.go
RUN go install /go/src/app/foo.go
ENTRYPOINT ["foo"]
