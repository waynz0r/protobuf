regenerate:
	go install github.com/waynz0r/protobuf/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=. --proto_path=../../../../../:../../protobuf/:. typedeclall.proto
