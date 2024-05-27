up:
	docker-compose -f docker-compose.yaml up --build

down:
	docker-compose -f docker-compose.yaml down

list:
	aws --endpoint-url=http://localhost:4566 s3 ls s3://my-bucket

create:
	aws --endpoint-url=http://localhost:4566 s3 mb s3://my-bucket

.PHONY: up down list create