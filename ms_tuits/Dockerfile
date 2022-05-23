FROM golang:1.17-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/ms_tuits

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/ms_tuits .

FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /tmp/ms_tuits/out/ms_tuits /app/ms_tuits


CMD ["/app/ms_tuits"]