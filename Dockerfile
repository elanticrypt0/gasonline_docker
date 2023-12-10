#-------------------------------------
# GOLANG
#-------------------------------------

FROM golang:alpine3.19 as go_dev
EXPOSE 3001
# GOLANG APP
WORKDIR /app/go_app

COPY ./go_app/. ./

RUN go mod tidy

# hot reload
RUN go install github.com/cosmtrek/air@latest

# Go app testing fase
FROM golang:alpine3.19 as go_testing

RUN go test test/.

# GOLANG APP
WORKDIR /app/go_app
COPY --from=go_dev ./app/go_app/. ./

RUN go test

# build binaries for amd64
FROM golang:alpine3.19 as build-amd64

# GOLANG APP
WORKDIR /var/bin/amd64
COPY --from=go_dev ./app/go_app/. /app/go_app

RUN rm -rf ./

RUN mkdir ./_db
RUN mkdir ./config
RUN mkdir ./seeds
RUN mkdir ./logs
RUN mkdir ./public

# build binaries for arm64
RUN cd ../../go_app && \
    GOOS=linux GOARCH=amd64 go build -o app -ldflags "-w -s" \
    && chmod +x app \
    && cd ../bin/amd64

COPY ./go_app/app ./

# COPY
COPY --from=go_dev /app/go_app/config/. ./config
COPY --from=go_dev /app/go_app/seeds/. ./seeds

#web user interface
COPY --from=wui_build /app/wui/dist/. ./public

FROM golang:alpine3.19 as build-arm

WORKDIR /var/bin/arm

RUN rm -rf ./.

RUN mkdir ./_db
RUN mkdir ./config
RUN mkdir ./seeds
RUN mkdir ./logs
RUN mkdir ./public

# create bin_arm
RUN cd ../../go_app && \
    GOOS=linux GOARCH=arm go build -o app -ldflags "-w -s" \
    && chmod +x app \
    && cd ../bin/amd64

COPY ./go_app/app ./

# COPY
COPY --from=go_dev ./go_app/config/. ./config
COPY --from=go_dev ./go_app/seeds/. ./seeds

#web user interface
COPY --from=wui_build /app/wui/dist/. ./public

# última etapa
FROM golang:alpine3.19 AS go_dev_runner
WORKDIR /app/go_app
COPY --from=go_dev ./. ./
# comando para iniciar la app
CMD [ "go","run","." ]

#-------------------------------------
# NODE + ASTRO + SVELTE + TAILWINDS
#-------------------------------------

# CREAR Web User Interface
FROM node:current-alpine3.18 as wui_dev

WORKDIR /app/wui

COPY ./wui/. ./

RUN npm install


# esta es una nueva etapa
FROM node:current-alpine3.18 AS wui_testing
WORKDIR /app/wui
# Esto copia los archivos de otro stage
COPY --from=wui_dev ./app/wui/. ./
COPY --from=wui_dev ./app/wui/node_modules/. ./node_modules

# realiza las pruebas automáticas
RUN npm run test

RUN rm -rf ./node_modules

# esta es una nueva etapa
FROM node:current-alpine3.18 AS wui_build
WORKDIR /app/wui
# Esto copia los archivos de otro stage
COPY --from=wui_testing ./app/wui/. ./.
COPY --from=wui_dev ./app/wui/node_modules/. ./node_modules

# crea el build
RUN npm run astro build

# última etapa
FROM node:current-alpine3.18 AS wui_dev_runner
EXPOSE 4321
WORKDIR /app/wui
COPY --from=wui_dev ./app/wui/. ./
COPY --from=wui_dev ./app/wui/node_modules/. ./node_modules

# comando para iniciar la app
CMD [ "npm","run","dev" ]

