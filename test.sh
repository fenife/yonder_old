
cur=`pwd`

# export GOPATH=$GOPATH:${cur}
export GOPATH=~/go:${cur}

echo $GOPATH

# go run main_test.go
go test main_test.go
