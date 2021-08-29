FROM golang:alpine AS builder

COPY . /

ENV CGO_ENABLED=0 \
  GOOS=linux

RUN apk --no-cache add git \
  && go build -a -installsuffix cgo -o /app kovercheng


FROM scratch

COPY --from=builder /app /

ENV DB_CONNECTURL="postgres://postgres:12345@host.docker.internal:5432/example"

ENTRYPOINT ["/app"]
