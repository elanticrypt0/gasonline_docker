#-------------------------------------
# GOLANG
#-------------------------------------

FROM golang:latest as go_dev

WORKDIR /app

# RUN apk install git
RUN apt-get update && apt-get install -y git
RUN git clone https://github.com/elanticrypt0/gasonline go_app

# GOLANG APP
COPY ./app/go_app/. ./

WORKDIR /app/go_app

RUN go mod tidy

# hot reload
# RUN go install github.com/cosmtrek/air@latest

# Go app testing fase
FROM golang:latest as go_testing

# GOLANG APP
WORKDIR /app/go_app
COPY --from=go_dev ./app/go_app/. ./
RUN go mod tidy

# RUN go test test/.
RUN go test

# build binaries for amd64
FROM golang:latest as build-amd64

# GOLANG APP
WORKDIR /app/bin/amd64
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

COPY ./app/go_app/app ./

# COPY
COPY --from=go_dev /app/go_app/config/. ./config
COPY --from=go_dev /app/go_app/seeds/. ./seeds

#web user interface
COPY --from=wui_build /app/wui/dist/. ./public

FROM golang:latest as build-arm

WORKDIR /app/bin/arm

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

COPY ./app/go_app/app ./

# COPY
COPY --from=go_dev ./app/go_app/config/. ./config
COPY --from=go_dev ./app/go_app/seeds/. ./seeds

#web user interface
COPY --from=wui_build /app/wui/dist/. ./public

# última etapa
FROM golang:latest AS go_dev_runner
EXPOSE 65065
WORKDIR /app/go_app
COPY --from=go_dev ./app/go_app/. ./
RUN go mod tidy

# comando para iniciar la app
CMD [ "go","run","." ]

#-------------------------------------
# NODE + ASTRO + SVELTE + TAILWINDS
#-------------------------------------

# CREAR Web User Interface
FROM node:latest as wui_dev

WORKDIR /app/wui

COPY ./app/wui/. ./

RUN npm install

# esta es una nueva etapa
FROM node:latest AS wui_testing
WORKDIR /app/wui
# Esto copia los archivos de otro stage
COPY --from=wui_dev ./app/wui/. ./


# realiza las pruebas automáticas
RUN npm run test

RUN rm -rf ./node_modules

# esta es una nueva etapa
FROM node:latest AS wui_build
WORKDIR /app/wui
# Esto copia los archivos de otro stage
COPY --from=wui_testing ./app/wui/. ./.


# crea el build
RUN npm run astro build

# última etapa
FROM node:latest AS wui_dev_runner
EXPOSE 4321
WORKDIR /app/wui
COPY --from=wui_dev ./app/wui/. ./

RUN npm install -g astro

# comando para iniciar la app
CMD [ "npm","run","start" ]

