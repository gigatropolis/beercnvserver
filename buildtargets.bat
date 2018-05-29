
set GOOS=windows
set GOARC=amd64
go build -o bin\beersvr_win_amd64\beersvr.exe beersvr.go

set GOOS=linux
set GOARC=amd64
go build -o bin\beersvr_linux_amd64\beersvr beersvr.go

set GOOS=linux
set GOARC=386
go build -o bin\beersvr_linux_386\beersvr beersvr.go

set GOOS=linux
set GOARC=arm
go build -o bin\beersvr_linux_arm\beersvr beersvr.go

set GOOS=linux
set GOARC=arm64
go build -o bin\beersvr_linux_arm64\beersvr beersvr.go

set GOOS=darwin
set GOARC=arm64
go build -o bin\beersvr_darwin_amd64\beersvr beersvr.go

set GOOS=windows
set GOARC=amd64

