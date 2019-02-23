
all: pre gen build

build:
	@echo "Build Client"
	@vgo build

pre:
	@echo "copy sample.proto to ..."
	@cp -f ../sample.proto proto/

gen:
	@echo "Generate grpc"
	@mkdir -p pb
	@protoc -I ./proto --go_out=plugins=grpc:./pb proto/sample.proto \
			-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			proto/sample.proto
clean:
	@rm -f proto/sample.proto
	@rm -rf pb
	@vgo clean
