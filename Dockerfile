FROM node:12.2.0-alpine as frontend
WORKDIR /src
ENV PATH /src/node_modules/.bin:$PATH
ENV NODE_ENV production

COPY frontend/ .
RUN \
    npm install --production && \
    npm install react-scripts -g && \
    npm run build

FROM golang:1.13-alpine as backend

RUN apk add git gcc build-base

# ENV CGO_ENABLED     0
ENV GOPROXY         https://proxy.golang.org
ENV GOOS            linux
# ENV GOARCH          amd64
# ENV GO111MODULE     on

WORKDIR /go/src/github.com/ramosisw/todo-go-reactjs
ADD . .
COPY --from=frontend /src/build/ frontend/build
RUN \
    go get github.com/go-bindata/go-bindata/go-bindata && \
    go get github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs
RUN go-bindata-assetfs -pkg=frontend -o=frontend/frontend.go -nocompress=false frontend/build/...

RUN go install -v -ldflags "-s -w"

FROM alpine as final
LABEL maintainer="Julio Ramos <ramos.isw@gmail.com>"
RUN apk add --no-cache ca-certificates
COPY --from=backend /go/bin/todo-go-reactjs /usr/bin/todo-go-reactjs

EXPOSE 80
VOLUME ["/data/"]
ENTRYPOINT todo-go-reactjs
