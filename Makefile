# Set up tools.
install:
	go install github.com/cosmtrek/air@v1.27.3

# Start dev server.
start:
	air

# Set up database.
init_db:
	./scripts/init_db.sh

# Down database.
down_db:
	./scripts/down_db.sh

migration_up:
	./scripts/migration_up.sh

migration_down:
	./scripts/migration_down.sh

# migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up