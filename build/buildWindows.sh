export GOOS=windows
export GOARCH=amd64
go build -ldflags="-s -w" -o ../tt.exe  ../cmd/time-tracker/main.go