cd docker || exit
docker compose up -d
docker compose cp ./postgres_data/sql postgres:/var/lib/postgresql/data
docker compose exec postgres psql -U postgres -d postgres -f /var/lib/postgresql/data/sql/reset_database.sql