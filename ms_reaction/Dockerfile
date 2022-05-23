FROM golang:1.17-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/ms_reaction

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/ms_reaction .

FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /tmp/ms_reaction/out/ms_reaction /app/ms_reaction


CMD ["/app/ms_reaction"]