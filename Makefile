.PHONY: run
run:
	go build && ./effective-go

.PHONY: migrate
migrate:
	cd sql/schema && goose postgres postgres://ussa:ussago@localhost:5432/go_blog up && cd ../..
	
.PHONY: down
down:
	cd sql/schema && goose postgres postgres://ussa:ussago@localhost:5432/go_blog down && cd ../..

.PHONY: generate
generate:
	sqlc generate

.PHONY: clean
clean:
	rm -f ./effective-go