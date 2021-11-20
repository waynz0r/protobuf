regenerate:
	go install github.com/waynz0r/protobuf/protoc-min-version
	go install github.com/waynz0r/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/waynz0r/protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/waynz0r/protobuf/types:. timefail.proto 
