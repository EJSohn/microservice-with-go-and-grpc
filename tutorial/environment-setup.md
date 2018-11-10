
# Environment Setup

전체 튜토리얼을 따라가기 위한 환경 설정을 합니다.

* go 설치
* grpc 서버를 만들기 위한 go 관련 패키지 설치
* grpc 웹 클라이언트를 만들기 위한 javascript 관련 패키지 설치

## go 설치

* Mac OS
```
brew install go
```

* Ubuntu
```
WIP
```

## grpc 서버를 만들기 위한 go 관련 패키지 설치

grpc에서 사용되는 Google protocol buffers의 go버젼을 설치하기위해서는 C++ protocol buffers 구현체가 필요합니다. 

다음은 [https://github.com/google/protobuf/releases](https://github.com/google/protobuf/releases)에서 최신 버젼의 protobuf-cpp zip를 받아 설치하는 과정입니다.

* Mac OS
```
brew install autoconf automake libtool

wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protobuf-cpp-<version>.zip

unzip protobuf-cpp-<version>.zip

cd protobuf-cpp-<version> && ./autogen.sh && ./configure && make

make check

sudo make install

# check installation
which protoc & protoc --version
```

* Ubuntu
```
WIP
```

Protocol Buffers를 설치하고 나면, 이를 go에서 사용하기 위해 필요한 패키지들과 grpc를 설치합니다.

* Mac OS
```
export GOBIN=${HOME}/go/bin
export PATH=${PATH}:${GOBIN}

go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

* Ubuntu
```
WIP
```