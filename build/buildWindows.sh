export GOOS=windows
export GOARCH=amd64
go build -ldflags="-s -w -X  'github.com/LReg/time-tracker/version.Version=0.0.1' -X '../cmd/time-tracker/main.go/version.Commit=$(git rev-parse HEAD)' -X 'main.Date=$(date -u +%Y-%m-%dT%H:%M:%SZ)'" -o ../tt.exe  ../cmd/time-tracker/main.go