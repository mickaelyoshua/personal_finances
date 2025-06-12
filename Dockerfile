FROM golang

WORKDIR /app

COPY ./app/go.mod .
COPY ./app/go.sum .

# Run and refresh server
RUN go install github.com/air-verse/air@latest
# Generate models and functions to interact with the database
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
# Generate templates for the frontend
RUN go install github.com/a-h/templ/cmd/templ@latest
# Do migrations in the database
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# Generate mocks for testing
RUN go install go.uber.org/mock/mockgen@latest

RUN go mod download
RUN go mod tidy