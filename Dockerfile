FROM golang:1.23.2-alpine AS apibuild
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o dist/main ./main.go

FROM golang:1.23.2-alpine AS api
WORKDIR /app
COPY --from=apibuild /app/dist/main .
EXPOSE 8090
CMD ["./main", "serve", "--http=0.0.0.0:8090"]

FROM node:20.18.1-alpine AS webbuild
WORKDIR /app
COPY client/package.json client/package-lock.json ./
RUN npm ci
COPY client/ ./
RUN npm run build
RUN npm prune --production

from node:20.18.1-alpine AS web
WORKDIR /app
COPY --from=webbuild /app/build ./build
COPY --from=webbuild /app/node_modules ./node_modules
COPY --from=webbuild /app/package.json ./package.json
EXPOSE 3000
CMD ["node", "build"]
