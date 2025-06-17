export GOOS=linux
export GOARCH=arm64
go build -ldflags="-s -w" -o ../tt-linux-arm64  ../cmd/time-tracker/main.go