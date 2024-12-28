# Usar una imagen base oficial de Go con Alpine Linux para reducir el tamaño de la imagen
FROM golang:1.21.3-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum al directorio de trabajo
COPY go.mod go.sum ./

# Descargar las dependencias del módulo
RUN go mod download

# Copiar el resto del código fuente al directorio de trabajo
COPY . .
ARG APP_PORT
ENV APP_PORT=$APP_PORT

# Copiar el archivo .env al directorio de trabajo
COPY .env .env

# Compilar el programa Go
RUN go build -o mi_programa 

# Exponer el puerto en el que la aplicación escuchará (usando la variable APP_PORT definida en .env)
EXPOSE $APP_PORT

# Especificar el comando para ejecutar el programa
CMD ["./mi_programa"]
