FROM golang:alpine AS build
RUN apk add --update gcc musl-dev 

WORKDIR /app
COPY . .
RUN go mod download
RUN go build .

FROM alpine
LABEL version=0.1.0
RUN apk add --no-cache --update procps
COPY --from=build /app/ps-exporter /bin/ps-exporter
ENTRYPOINT  [ "/bin/ps-exporter" ]

EXPOSE 9095
