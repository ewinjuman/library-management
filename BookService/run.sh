export DB_HOST="localhost"
export DB_USER="postgres"
export DB_PASSWORD="12345"
export DB_NAME="library"
export DB_PORT="5432"
export REDIS_ADDRESS="123"
export HTTP_PORT="8080"
export GRPC_PORT="5080"
export JWT_SECRET="secret"
export AUTHOR_GRPC="localhost:5083"
export CATEGORY_GRPC="localhost:5084"
export USER_GRPC="localhost:5081"

go run main.go