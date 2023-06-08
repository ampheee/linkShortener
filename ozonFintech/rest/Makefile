
app-on-redis-build:
	docker compose build --build-arg STORAGE=Redis

app-on-postgres-build:
	docker compose build --build-arg STORAGE=PostgreSQL --build-arg MIGRATION=OnStart

app-on-redis-start:
	docker-compose up -d app redis

app-on-postgres-start:
	docker-compose up -d app postgres