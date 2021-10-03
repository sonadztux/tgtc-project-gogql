# run gql service
go_run_gql:
	@go build -v -o bin/gql_service gql/main.go
	@./bin/gql_service