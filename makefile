
# คำสั่ง generate api spectification
codegen:
	oapi-codegen -package=http -generate=types,fiber -o library/http/api.gen.go api/openapi.yaml	

.PHONY: codegen
