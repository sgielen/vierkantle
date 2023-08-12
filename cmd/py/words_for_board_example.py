from pyease_grpc import RpcSession, RpcUri
import json
from base64 import b64encode, b64decode

board = {
	"width": 3,
	"height": 3,
	"cells": [
		[ 'v', 'e', 'r' ],
		[ 'l', 'e', 't' ],
		[ 'l', 'e', 'r' ],
	],
}
boardJson = json.dumps(board).encode('ascii')

endpoint = "https://vierkantle.nl/api"
session = RpcSession.from_file("../../pkg/proto/vierkantle.proto")
uri = RpcUri(
	base_url=endpoint,
	package="nl.vierkantle",
	service="VierkantleService",
	method="WordsForBoard",
)
response = session.request(uri, {"board": b64encode(boardJson)})
print("Score: ", response.single['score'])
boardJson = b64decode(response.single['board'])
board = json.loads(boardJson)

print(json.dumps(board, indent=2))
