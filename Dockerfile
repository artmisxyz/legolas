FROM golang:1.17 AS build

WORKDIR /app
COPY . ./
RUN go build --ldflags "-linkmode external -extldflags '-static'" -v -o ./legolas .

FROM alpine:latest
ARG BUILD_DATE
ARG VCS_REF
ARG BUILD_VERSION

RUN apk --no-cache --update add ca-certificates

RUN mkdir /app

COPY --from=build ./app/legolas /app

WORKDIR /app
CMD ["/app/legolas"]
