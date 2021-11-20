regenerate:
	go install github.com/waynz0r/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/waynz0r/protobuf/protobuf/ \
	--gogo_out=. events.proto 
