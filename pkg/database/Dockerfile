FROM amacneil/dbmate:2.4.0

ENV DBMATE_MIGRATIONS_TABLE=vierkantle.schema_migrations

COPY migrations db/migrations

ENTRYPOINT [""]
CMD ["dbmate", "--wait", "status"]
