#stage 1
FROM node:16 as node
WORKDIR /app
COPY . .
# RUN npm install
RUN npm run build --prod --openssl-legacy-provider
#stage 2
FROM nginx:latest
COPY --from=node /app/dist/front-end /usr/share/nginx/html
