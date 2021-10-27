FROM golang:alpine AS builder

WORKDIR /
COPY . .

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GO111MODULE=on

RUN go build -p 4 -a -installsuffix cgp -o cmd/app kovercheng


FROM scratch

COPY --from=builder /cmd .

ENV SERVER_URL="0.0.0.0:3000" \
    PGCONNECTURL="postgres://postgres:12345@host.docker.internal:5432/examplePg" \
    MONGOCONNECTURL="mongodb://mongo:12345@host.docker.internal:27017/" \
    REDISCONNECTURL="redis://:12345@host.docker.internal:6379/0"

EXPOSE 3000
ENTRYPOINT ["/app"]
