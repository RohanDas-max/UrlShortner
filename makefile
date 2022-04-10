run:
	docker-compose up -d --build
serve:
	go run cmd/server/server.go
