FROM golang:alpine as builder

LABEL maintainer="Ralf Geschke <ralf@kuerbis.org>"
LABEL last_changed="2022-07-26"

RUN apk update && apk add --no-cache git

# Build invafetch
RUN mkdir /build
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o invafetch .

# Build minimal image with dynpower and dynpower-cli only
FROM scratch
COPY --from=builder /build/invafetch /app/

ENV PATH "$PATH:/app"
WORKDIR /app
CMD ["./invafetch","start"]
