# Usa la imagen base oficial de Python 3.12
FROM python:3.12.3

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el archivo de requirements a la imagen
COPY requirements.txt .

# Instala las dependencias especificadas en requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Copia el resto del código de la aplicación a la imagen
COPY . .

# Configurar variables de entorno usando build args
ARG APP_HOST
ARG APP_PORT
ARG MYSQL_HOST
ARG MYSQL_USER
ARG MYSQL_PASSWORD
ARG MYSQL_DB

ENV APP_HOST=$APP_HOST
ENV APP_PORT=$APP_PORT
ENV MYSQL_HOST=$MYSQL_HOST
ENV MYSQL_USER=$MYSQL_USER
ENV MYSQL_PASSWORD=$MYSQL_PASSWORD
ENV MYSQL_DB=$MYSQL_DB

# Expone el puerto en el que correrá la aplicación (cambia según sea necesario)
EXPOSE 9091

# Comando por defecto para correr la aplicación (ajusta según sea necesario)
CMD ["python", "index.py"]
