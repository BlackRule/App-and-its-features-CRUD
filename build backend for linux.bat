set GOOS=linux
set GOARCH=amd64
set GOPATH=%cd%\gopath
cd src/backend
rem go install
go build -o "..\..\dst\exe" ".\cmd\App-and-its-features-CRUD\App-and-its-features-CRUD.go"
cd ..\..