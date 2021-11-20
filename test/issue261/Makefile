regenerate:
	go install github.com/waynz0r/protobuf/protoc-min-version
	go install github.com/waynz0r/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --gogoslick_out=\
	Mgoogle/protobuf/duration.proto=github.com/waynz0r/protobuf/types,\
	:. \
	--proto_path=../../../../../:../../protobuf/:. issue261.proto
