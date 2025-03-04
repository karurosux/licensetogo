FROM golang:1.23.2-alpine AS api
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o dist/main ./main.go
RUN mkdir /licensetogo \
&& cp ./dist/main /licensetogo
WORKDIR /
RUN rm -rf /app
WORKDIR /licensetogo
EXPOSE 8090
CMD ["./main", "serve", "--http=0.0.0.0:8090"]

FROM node:20.18.1-alpine AS web
WORKDIR /app
COPY client/package.json client/package-lock.json ./
RUN npm ci
COPY client/ ./
RUN npm run build
RUN npm prune --production
RUN mkdir /licensetogo \
&& cp -r ./build /licensetogo/build \
&& cp -r ./node_modules /licensetogo/node_modules \
&& cp ./package.json /licensetogo/package.json
WORKDIR /
RUN rm -rf /app
WORKDIR /licensetogo
EXPOSE 3000
CMD ["node", "build"]

