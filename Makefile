stack-up:
	docker-compose -f deploy/local/docker-compose.yaml  up -d --build

stack-down:
	docker-compose -f deploy/local/docker-compose.yaml down

bulb-color:
	@echo $$(curl -s http://localhost:3333/echo)

mock:
	mockery

testing:
	go test ./... -cover