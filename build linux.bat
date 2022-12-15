set GOOS=linux
set GOARCH=amd64
cd src
rem go install
go build -o "..\dst\exe" ".\main.go"
cd ..