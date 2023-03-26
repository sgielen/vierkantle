Vierkantle
----------

A Dutch Squardle clone.

To generate a board offline:

```
go run ./cmd/generator contrib/opentaal-wordlist/wordlist.txt
```

To run the backend locally, first run PostgreSQL:

```
$ docker-compose -f docker-compose.infrastructure.yml up -d
$ export VIERKANTLE_DSN="postgres://vierkantle:vierkantle@localhost:5433/vierkantle"
```

You can now connect and inspect the database using `psql $VIERKANTLE_DSN`.

Import the schema before continuing:

```
cat pkg/database/schema.sql | psql "$VIERKANTLE_DSN"
```

Then, run the backend:

```
gow -r=false run github.com/sgielen/vierkantle/cmd/backend -words contrib/opentaal-wordlist/wordlist.txt -bonus contrib/SUBTLEX-NL.cd-above2.txt
```

To run the frontend:

```
cd cmd/ui && yarn serve
```

The port number used will be in the `yarn serve` output.
