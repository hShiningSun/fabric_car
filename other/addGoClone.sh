git clone https://github.com/golang/sys.git    $GOPATH/src/golang.org/x/sys
git clone https://github.com/grpc/grpc-go.git  $GOPATH/src/google.golang.org/grpc
git clone https://github.com/golang/net.git    $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git   $GOPATH/src/golang.org/x/text
git clone https://github.com/golang/crypto.git $GOPATH/src/golang.org/x/crypto
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
echo "cd $GOPATH/src/"
echo "go install google.golang.org/grpc"