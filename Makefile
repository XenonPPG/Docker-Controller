generate-doc:
	docker-compose build protodoc
	docker-compose run --rm protodoc
	docker-compose down --rmi local

generate-proto:
	docker-compose build protogen protodoc
	docker-compose run --rm protogen
	docker-compose run --rm protodoc
	docker-compose down --rmi local

.PHONY: generate-doc, generate-proto