FROM golang:1.15 AS builder

WORKDIR $GOPATH/src/health

COPY . ./

RUN go get -u github.com/swaggo/swag/cmd/swag@v1.6.7
RUN swag init

RUN go get -u
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:3.12.3

COPY --from=builder /go/src/health/main ./
COPY --from=builder /go/src/health/configs ./configs

EXPOSE 8080

ENTRYPOINT ["./main"]
