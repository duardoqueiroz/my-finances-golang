# Caminho para o arquivo config.json
CONFIG_FILE="pkg/config/config.json"

# Verifica se o arquivo existe
if [ ! -f "$CONFIG_FILE" ]; then
    echo "Erro: Arquivo config.json não encontrado."
    exit 1
fi

# Extrai as informações do arquivo JSON usando jq (precisa estar instalado)
DB_HOST=$(jq -r '.database.host' "$CONFIG_FILE")
DB_PORT=$(jq -r '.database.port' "$CONFIG_FILE")
DB_USER=$(jq -r '.database.user' "$CONFIG_FILE")
DB_PASSWORD=$(jq -r '.database.password' "$CONFIG_FILE")
DB_NAME=$(jq -r '.database.name' "$CONFIG_FILE")

# Constrói a string de conexão
DB_CONNECTION="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"

# Executa o comando de migração
migrate -database "$DB_CONNECTION" -path db/migrations down