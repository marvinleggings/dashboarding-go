docker run --rm -p 8080:80 -ti -v "$PWD"/:/go/src/ --name dreamfrontend golang:1.12.1-alpine3.9 go run /go/src/app/main.go || docker restart dreamfrontend
