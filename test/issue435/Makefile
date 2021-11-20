regenerate:
	go install github.com/waynz0r/protobuf/protoc-gen-gogo
	protoc --gogo_out=. --proto_path=../../../../../:../../protobuf/:. *.proto
