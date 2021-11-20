regenerate:
	go install github.com/waynz0r/protobuf/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/waynz0r/protobuf/types:. --proto_path=../../../../../:../../protobuf/:. issue260.proto
