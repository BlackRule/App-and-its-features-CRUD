set GOOS=windows
set GOARCH=amd64
set GOPATH=%cd%\gopath
cd src/backend
REM go install
go build -o "..\..\dst\exe.exe" ".\main.go"
cd ..\..
