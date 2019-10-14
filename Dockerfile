FROM node:12.2.0-alpine as frontend
WORKDIR /src
ENV PATH /src/node_modules/.bin:$PATH

COPY frontend/ .
RUN \
    npm install --production && \
    npm install react-scripts -g && \
    npm run build

FROM golang:1.13-alpine as backend
LABEL maintainer="Julio Ramos <ramos.isw@gmail.com>"
RUN apk add git gcc build-base

# ENV CGO_ENABLED     0
ENV GOPROXY         https://proxy.golang.org
ENV GOOS            linux
# ENV GOARCH          amd64
# ENV GO111MODULE     on

WORKDIR /go/src/github.com/jramos/golang-reactjs
ADD . .
RUN go install -v -ldflags "-s -w"

FROM alpine as final
RUN apk add --no-cache ca-certificates
ENV PATH_FRONTEND /usr/bin/static
COPY --from=frontend /src/build/ /usr/bin/static
COPY --from=backend /go/bin/golang-reactjs /usr/bin/golang-reactjs

EXPOSE 80
VOLUME ["/data/"]
ENTRYPOINT golang-reactjs