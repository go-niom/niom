package misc

const MiscDockerFile = `
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

const MiscDockerIgnore = `
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
