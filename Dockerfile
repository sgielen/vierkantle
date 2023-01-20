FROM golang:1.19.3-buster AS build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN cd cmd/backend && go build -o /backend

FROM debian:buster-20221114 AS base
RUN apt-get update && apt-get --no-install-recommends install -y \
	wget ca-certificates \
	&& rm -rf /var/lib/apt/lists/*

WORKDIR /
EXPOSE 80

FROM base AS backend
COPY contrib/SUBTLEX-NL.cd-above2.txt /data/wordlist.txt
COPY contrib/opentaal-wordlist/wordlist.txt /data/bonuslist.txt
COPY --from=build /backend /
USER nobody:nogroup
ENTRYPOINT [ "/backend" ]
