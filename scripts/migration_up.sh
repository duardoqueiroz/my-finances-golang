# Carregar variáveis de ambiente do arquivo .env
. ./.env

# Construir a string de conexão
CONNECTION_STRING="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable"

# Executa o comando de migração
migrate -database "$CONNECTION_STRING" -path db/migrations up