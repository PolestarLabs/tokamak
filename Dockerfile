FROM golang:alpine AS builder
WORKDIR /src
COPY . /src
RUN cd ./src && go build -o tokamak

FROM alpine
WORKDIR /app
COPY --from=builder /src/src/tokamak /app

EXPOSE 1234
ENTRYPOINT ./tokamak
