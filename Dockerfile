FROM golang:1.16-rc-alpine as builder

ENV GO111MODULE on
ENV GOPROXY https://proxy.golang.org/

# Installing nodejs and other dependecies
RUN apk add --update nodejs-current python2 curl bash build-base

# Installing Yarn
RUN curl -o- -L https://yarnpkg.com/install.sh | bash -s -- --version 1.22.10
ENV PATH="$PATH:/root/.yarn/bin:/root/.config/yarn/global/node_modules"

WORKDIR /sounio_health
ADD go.mod .
ADD go.sum .
RUN go mod download -x

# Installing ox
RUN go install github.com/wawandco/oxpecker/cmd/ox@master
ADD . .

# Building the application binary in bin/app
RUN ox build --static -o bin/app

FROM alpine

# Binaries
COPY --from=builder /sounio_health/bin/app /bin/app
# COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /zoneinfo.zip
# ENV ZONEINFO=/zoneinfo.zip

EXPOSE 3000

# For migrations use
# CMD ox db migrate up; app
CMD app