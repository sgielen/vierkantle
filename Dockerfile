FROM golang:1.23.1-bullseye AS build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN cd cmd/backend && go build -o /backend
RUN cd cmd/board && go build -o /board

FROM debian:bullseye-20230919 AS base
RUN apt-get update && apt-get --no-install-recommends install -y \
	wget ca-certificates \
	&& rm -rf /var/lib/apt/lists/*

WORKDIR /
EXPOSE 80

FROM base AS backend
COPY contrib/SUBTLEX-NL.cd-above2.txt /data/wordlist.txt
COPY contrib/opentaal-wordlist/wordlist.txt /data/bonuslist.txt
COPY --from=build /backend /
COPY --from=build /board /bin
USER nobody:nogroup
ENTRYPOINT [ "/backend" ]
