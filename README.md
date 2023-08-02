# Start
Run `docker compose -f docker-compose-redis.yml up` to start redis cache locally. If cache can not be reached the errors will be overlooked here and request will be make against ERPLY API.

Once redis is running start the api by running `go run .`

# Swagger doc
Swagger URL is localhost:5123/swagger/index.html

# Unit tests

To run unit tests run `go test ./...`
