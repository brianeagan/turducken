build:
	GOOS=js GOARCH=wasm go build -o bin/turducken.wasm turducken.go

serve:
	go run MEGASERVER.go

open:
	open 'http://localhost:8080/turducken.html'