set GOOS=windows
set GOARCH=amd64
set GOPATH=%cd%\gopath
cd src/backend
rem go install
go build -o "..\..\dst\exe.exe" ".\main.go"
cd ..\..
