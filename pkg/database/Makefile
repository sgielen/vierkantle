MIGRATION_FILES = $(shell find migrations -type f -name '*')
QUERY_FILES = $(shell find . -type f -name 'queries*.sql')

gendb/queries.sql.go: sqlc.yaml $(MIGRATION_FILES) $(QUERY_FILES)
	dbmate dump
	sqlc generate
