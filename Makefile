gen:
	cd api; buf generate

run:
	go run cmd/main.go

clean:
	rm -rf pkg