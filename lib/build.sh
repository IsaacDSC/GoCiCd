env GOOS=linux GOARCH=amd64 go build -o ./dist/linux-amd64 ./lib/main.go
env GOOS=linux GOARCH=arm64 go build -o ./dist/linux-arm64 ./lib/main.go
env GOOS=windows GOARCH=amd64 go build -o ./dist/windows-amd64 ./lib/main.go
env GOOS=windows GOARCH=386 go build -o ./dist/windows-386 ./lib/main.go