set GOOS=linux
set GOARCH=amd64
set GOPATH=%cd%\dst
cd src
rem go install
go build
cd ..