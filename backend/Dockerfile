# ----
# base
# ----
FROM golang:1.17-alpine3.14 as base

ENV GOOS=linux
ENV CGO_ENABLED=0

WORKDIR /go/src/

COPY ./go.sum /go/src/go.sum
COPY ./go.mod /go/src/go.mod

RUN go mod download

# ----
# dev
# ----
FROM base as dev

# -------
# builder
# -------
FROM base as builder

COPY ./ /go/src/

RUN go build -o /go/bin/api cmd/api/main.go

# ----------
# production
# ----------
FROM scratch as production

COPY --from=builder /go/bin/api /go/bin/api
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/go/bin/api"]
