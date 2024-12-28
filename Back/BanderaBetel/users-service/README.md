# Generator Python

Proyecto generado para la configuración automática de microservicios con python

## Requisitos

* [Instalar NodeJs LTS](https://nodejs.org/es/)
* Instalar Yeoman
    `npm install -g yo`
* Descargar dependencias 
    `npm install`
* Ejecutar dentro de generator-python/
    `npm link`

## Instalar proyecto

- Ejecutar el comando: `yo python`  

- En el archivo *config.json* que se encuentra en la carpeta *generators/app* se configuran los parametros para el nombre de la compañia, nombre del proyecto y puerto

```json
{
  "company": "Axity",
  "name": "Userss",
  "port": 9090
}
```

- Esto genera un proyecto llamado: [nombre ingresado]-service

## Comandos

1. **Crear entorno virtual:**
    `python -m venv env`

2. **Activar entorno virtual:**
    `env\Scripts\activate`

3. **Instalar los paquetes del archivo requirements.tx, requeridos en nuestro entorno virtual:**
    `pip install -r ./requirements.txt`

4. **En listar los paquetes instalados en el entorno virtual:**
    `pip list`

5. **Ejecutar servicio:**
    `python ./index.py`

6. **Ejecutar pruebas unitarias**
    `pytest TestProyectFacade.py`

7. **Guarda todos los paquetes instalos en nuestro entorno virtual en requirements.txt**
    `pip freeze > requirements.txt`

8. **Desactivar el entorno virtual:**
    `deactivate`


    ## Docker

    Docker automatiza el despliegue de aplicaciones dentro de contenedores.
    
    El archivo Dockerfile contiene las instrucciones para crear la imagen.
    
    Pasos para desplegar la aplicación:
    1. **Generar imagen**
    ```sh
    docker build -t my-python-app --build-arg APP_HOST=0.0.0.0 --build-arg APP_PORT=9091 --build-arg MYSQL_HOST=111.111.1.11 --build-arg MYSQL_USER=root --build-arg MYSQL_PASSWORD=root --build-arg MYSQL_DB=dataBase .
    ```
    2. **Consultar imagenes**
    ```sh
    docker images
    ```
    3. **Ejecutar imagen**
    ```sh
    docker run --name my-app -p 9091:9091 my-python-app
    ```
    4. **Mostrar contenedores en ejecución**
    ```sh
    docker ps
    ```
    5. **Obtener los registros de un contenedor**
    ```sh
    docker container logs [containerid]
    ```
    6. **Aplicación desplegada**
    - Validar que la aplicación se este ejecutando por el puerto 9091 http://127.0.0.1:9091/swagger/index.html

## Configuración

1. **Configurar variables de entorno**  
En el archivo **config.py** configure las variables de entorno para su proyecto.
```python
class Config:
    DEBUG = False
    SECRET_KEY = "secret"
    # DB Credentials
    MYSQL_HOST = "localhost" # Host de la base de datos
    MYSQL_USER = "root" # Usuario de la base de datos
    MYSQL_PASSWORD = "root" # Contraseña de la base de datos
    MYSQL_DB = "db_python_test" # Nombre de la base de datos
    # SQL ALCHEMY CONNECTION
    SQLALCHEMY_TRACK_MODIFICATIONS = False
    SQLALCHEMY_DATABASE_URI = (
        f"mysql+mysqlconnector://{MYSQL_USER}:{MYSQL_PASSWORD}@{MYSQL_HOST}/{MYSQL_DB}"
    )
```
## ¿Cómo usarlo?

1. **Operaciones a la base de datos**
  - Consulta de datos por llave primaria
```python
def GetByIdAsync(self, id: int) -> ProyectModel:
        return db.session.get(ProyectModel, id)
```
  - Consulta de una lista de datos
```python
def GetAllAsync(self) -> list[ProyectModel]:
        return db.session.query(ProyectModel).filter(ProyectModel.Active == True).all()
```
  - Inserción a la base de datos
```python
def InsertAsync(self, model: ProyectModel) -> bool:
        db.session.add(model)
        db.session.commit()
        return True
```
- Actualización a la base de datos
```python
def UpdateAsync(self, model: ProyectModel) -> bool:
        db.session.commit()
        return True
```

## Capas del proyecto
1. **Api:** Encontramos los controladores y la configuración principal.
2. **Facade:** Es el orquestador de servicios, en esta capa se utiliza la menor lógica posible y solo se realizan llamadas a la capa de servicio.
3. **Services:** Integración de servicios a terceros, reglas del negocio y llamadas a los Data Access Object (DAO).
4. **Persistence:** Encontramos los Data Access Object (DAO), son las clases que interactúan con la base de datos.
5. **Model:** Marco de mapeo relacional para acceder a base de datos.
6. **Common:** Encontramos los Data Transfer Object (Dtos), Excepciones, Enums y Funciones de validación que puedan ser reutilizados en todo el proyecto, con el proposito de reducir la complejidad ciclomática.
7. **Test:** Capa para las pruebas unitarias (TDD).

## Colaboradores

**Jonathan Aldana Herrera**  
*[jonathan.aldana@axity.com]*   
**Ingrid Mendoza Cabrera**  
*[ingrid.mendoza@axity.com]*  

![](./assets/CReA.png)

## Licencia

[MIT](https://opensource.org/licenses/MIT)