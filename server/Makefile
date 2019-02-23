
all: pre gen build

build:
	@echo "Build Server"
	@vgo build

pre:
	@echo "copy sample.proto to ..."
	@cp -f ../sample.proto proto/

gen:
	@echo "Generate grpc"
	@mkdir -p pb
	@protoc -I ./proto \
			-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc:./pb \
			--grpc-gateway_out=logtostderr=true:./pb \
			--swagger_out=logtostderr=true:./pb \
			proto/sample.proto
clean:
	@rm -f proto/sample.proto
	@rm -rf pb
	@vgo clean
