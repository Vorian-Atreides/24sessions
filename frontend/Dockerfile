FROM node:10.0 AS builder
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

FROM nginx:1.13-alpine
COPY --from=builder /app/build /usr/share/nginx/html
