install-deps:
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger.yaml:
	swagger generate spec -o ./swagger.yaml

serve-swagger: swagger.yaml
	swagger serve -F=swagger swagger.yaml --no-open