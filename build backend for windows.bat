set GOOS=windows
set GOARCH=amd64
set GOPATH=%cd%\gopath
cd src/backend
REM go install
go build -o "..\..\dst\exe.exe" ".\cmd\App-and-its-features-CRUD\App-and-its-features-CRUD.go"
cd ..\..
