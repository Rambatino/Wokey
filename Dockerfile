FROM golang:latest
WORKDIR /go/src/wokey
COPY main.go glide.yaml glide.lock server.go client.go hub.go ./
COPY app app/
COPY database database/
COPY routes routes/
RUN curl https://glide.sh/get | sh
RUN glide install
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.8
WORKDIR /root/
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/wokey/main .
COPY public public/
COPY src src/
COPY package.json package-lock.json tsconfig.json ./
RUN apk add --update nodejs nodejs-npm
RUN npm install
RUN npm run-script build
COPY .env .env
CMD ["./main"]
