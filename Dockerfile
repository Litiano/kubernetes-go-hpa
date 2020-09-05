FROM golang:1.13 as builder

WORKDIR /go/src/app

#disable crosscompiling
ENV CGO_ENABLED=0
#compile linux only
ENV GOOS=linux
#only amd64
ENV GOARCH=amd64

COPY src/main.go main.go

RUN go build -ldflags '-w -s' -a -installsuffix cgo -o main.o

FROM scratch
COPY --from=builder /go/src/app/main.o /main.o
EXPOSE 80
CMD ["/main.o"]