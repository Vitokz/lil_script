FROM golang:alpine AS builder

# Set up dependencies
ENV PACKAGES git build-base

# Set up workdir path
ENV APP_HOME /go/src/bs_helper

WORKDIR "$APP_HOME"

# Install dependencies
RUN apk add --update $PACKAGES
RUN apk add linux-headers
#RUN apt-get update && apt-get install build-essential -y

COPY . .

# build binary
RUN go mod download
RUN go mod verify
RUN go build -o /go/src/bs_helper/bs_helper

# Final image
FROM alpine

WORKDIR /
COPY --from=builder /go/src/bs_helper/bs_helper /usr/bin/bs_helper

CMD ["bs_helper"]