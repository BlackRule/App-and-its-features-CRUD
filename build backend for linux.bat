set GOOS=linux
set GOARCH=amd64
set GOPATH=%cd%\gopath
cd src/backend
rem go install
go build -o "..\..\dst\exe" ".\main.go"
cd ..\..