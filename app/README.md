# Gasonline

GASONLINE es un SDK que incluye Go + Astro + Svelte (GAS)
Ideal para crear APIs

## Diseño RFM (Routes Features Models)

### Routes

Las rutas se definen en la carpeta routes dentro del archivo setuproutes allí se agregan las funciones. Para poder acceder a las funcionalidades de Fiber o Gorm o la configuración de la app se debe pasar una variable

    gasonline *webcore.GasonlineApp

### Features

Las "features" son los controladores de las rutas y reciben 2 parámetros: 1) el contexto de Fiber y otra variable de *webcore.gasonlineApp para poder acceder a la db y configuración.

### Models

Los modelos deben recibir también la variable que apunte a *webcore.gasonlineApp

## Public

los archivos publicos están dentro de la carpeta public y por defecto tienen varios accesos. Esto se puede ver en webcore_features/routes.go

# Astro + Svelte

Para la creación del front-end, wui (Web User Interface). En linux ejecutar el script *create_wui.sh*

Esto crea una carpeta _wui_, allí se encuentran los archivos para realizar la interfaz gráfica. Allí se puede trabajar en modo desarrollo, una vez terminado se puede compilar con el comando

    npm run astro build

# Generar un build

Para generar un build se debe ejecutar el script _build.sh_ eso hará varias cosas:

1- Borrá la carpeta _build_ si está creada.
2- Creará los binarios para windows y linux
3- Copiará la carpeta generada por el _wui_
4- copiará los archivos de configuración