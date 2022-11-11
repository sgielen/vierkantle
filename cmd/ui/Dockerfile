FROM node:18.12 AS build
WORKDIR /app
COPY . .
RUN yarn install --frozen-lockfile && yarn quasar build && du -hs dist

FROM nginx:1.21.6
COPY --from=build /app/dist/spa /usr/share/nginx/html
COPY nginx-default.conf /etc/nginx/conf.d/default.conf
