package utils

const Env = `
# APP settings:
APP_NAME={{ .NameLowerCase}}
APP_HOST="0.0.0.0"
APP_PORT=4000
APP_READ_TIMEOUT=30
APP_DEBUG=false

# JWT settings:
JWT_SECRET_KEY="super_secret_here"
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=1440

# Database settings:
DB_HOST=niom-postgres
DB_PORT=5432
DB_USER=dev
DB_PASSWORD=dev
DB_NAME=niom_go_api
DB_SSL_MODE=disable
DB_DEBUG=true
DB_MAX_OPEN_CONNECTIONS=3
DB_MAX_IDLE_CONNECTIONS=1
DB_MAX_LIFETIME_CONNECTIONS=10
`

const DockerFile = `
# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
ENV GOPATH /go
WORKDIR /go/src
COPY . /go/src/app-build
RUN cd /go/src/app-build && env GOOS=linux CGO_ENABLED=0 go build .

#final stage
FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=builder /go/src/app-build/{{ .NameLowerCase}} /app
COPY .env /app

LABEL Name={{ .NameLowerCase}} Version=0.0.1
EXPOSE 4000
CMD ["./{{ .NameLowerCase}}"]
`

const DockerIgnore = `
**/.classpath
**/.dockerignore

**/.git
**/.gitignore
**/.project
**/.settings
**/.toolstarget
**/.vs
**/.vscode
**/*.*proj.user
**/*.dbmdl
**/*.jfm
**/bin
**/charts
**/docker-compose*
**/compose*
**/Dockerfile*
**/node_modules
**/npm-debug.log
**/obj
**/secrets.dev.yaml
**/values.dev.yaml
README.md

`

const MainGo = `
package main

import (
	// load API Docs files (Swagger)

	"{{ .ModuleName}}/server"
	// _ "{{ .ModuleName}}/docs"
	"{{ .ModuleName}}/pkg/config"
)

// @title Travel App
// @version 1.0
// @description Travel App Backend REST API
// @in header
// @name Authorization
// @host localhost:3000
// @BasePath /api
func main() {

	// setup various configuration for app
	config.LoadAllConfigs(".env")
	server.Serve()
}


`
