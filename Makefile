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