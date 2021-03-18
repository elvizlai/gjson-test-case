pb:
	@protoc -I=. --experimental_allow_proto3_optional \
		--go_out=. --go_opt=paths=source_relative \
		--go-json_out=. \
		common.proto
