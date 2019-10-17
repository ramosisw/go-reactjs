# todo-go-reactjs
Todo application with Go as Backend and ReactJS as Frontend

![Todo List](https://raw.githubusercontent.com/ramosisw/todo-go-reactjs/master/screenshots/todo-list.png)

## How to build frontend
The following Go dependencies are required to compile

```sh
go get github.com/go-bindata/go-bindata/go-bindata
go get github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs
```

The next thing is to compile the frontend for that you have to have nodejs and execute the following commands

```sh
npm --prefix=frontend install
npm --prefix=frontend run build
```

Now we compile the assets of the build folder

```sh
go-bindata-assetfs -pkg=frontend -nocompress=false -o=frontend/frontend.go frontend/build/...
```

## Build binary
After following the previous steps we execute the following command.
```sh
go build -v -ldflags "-s -w"
# Ejecute
todo-go-reactjs
```

## Build with go script
Run
```sh
go run ./gobuild
```

Example
```bash
2019/10/16 16:39:12 Building please wait...
2019/10/16 16:39:12 Runing: go get github.com/go-bindata/go-bindata/go-bindata
2019/10/16 16:39:13 Runing: go get github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs
2019/10/16 16:39:14 Runing: npm install
2019/10/16 16:39:38 Runing: npm run build
2019/10/16 16:39:52 Runing: go-bindata-assetfs -pkg=frontend -nocompress=false -o=frontend/frontend.go frontend/build/...
2019/10/16 16:39:52 Runing: go build -v -ldflags -s -w
``` 

## Run from docker
```sh
docker run --rm ramosisw/todo-go-reactjs -p 8080:80
```
and open http://localhost:8080 if you use docker-machie probably hosted on http://192.168.99.10:8080
