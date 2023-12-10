# Resumen

La imagen de docker crear una carpeta para la app de gasoline e instala las dependecias.
Agrega air para poder utilizar el hot-reload.

También la imágen crea las dependecias para crear la Web User Interface (wui) instalando node con astro, svelte, tailwinds y flowbite.



# Crear imagen docker

Ejecutar este comando para crear una imagen de docker

    docker build -t go_docker:1.0.0 .

## Para crear un binario específico

    docker build --target=amd64

# Correr una imágen

    docker container run [nombre]

## Para correr un binario específico
