APP_NAME=$1

env GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o bin/${APP_NAME}/bootstrap app/${APP_NAME}/main.go
zip -j bin/${APP_NAME}.zip bin/${APP_NAME}/bootstrap
