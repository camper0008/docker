# frontend

FROM node:16

COPY src /docker/frontend

WORKDIR /docker/frontend

RUN npm ci --only=production
RUN npm run build
