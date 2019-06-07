FROM golang:1.11.5
WORKDIR /go/src/wokey
COPY main.go glide.yaml glide.lock server.go ./
COPY app app/
COPY routes routes/
RUN curl https://glide.sh/get | sh
RUN glide install
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.8
WORKDIR /root/
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/wokey/main .
COPY public public/
COPY build build/
COPY .env .env
CMD ["./main"]
