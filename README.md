Vierkantle
----------

A Dutch Squardle clone.

To generate a board offline:

```
go run ./cmd/generator contrib/opentaal-wordlist/wordlist.txt
```

To run the backend:

```
gow -r=false run github.com/sgielen/vierkantle/cmd/backend -words contrib/opentaal-wordlist/wordlist.txt -bonus contrib/SUBTLEX-NL.cd-above2.txt
```

To run the frontend:

```
cd cmd/ui && yarn serve
```

The port number used will be in the `yarn serve` output.
