from decouple import config as configEnv
class ConfigBase:
    DEBUG = False
    SECRET_KEY = "secret"
    # Host to expose
    # DB Credentials
    MYSQL_HOST = configEnv('MYSQL_HOST')
    MYSQL_USER = configEnv('MYSQL_USER')
    MYSQL_PASSWORD = configEnv('MYSQL_PASSWORD')
    MYSQL_DB = configEnv('MYSQL_DB')
    # SQL ALCHEMY CONNECTION
    SQLALCHEMY_TRACK_MODIFICATIONS = False
    SQLALCHEMY_DATABASE_URI = (
        f"mysql+mysqlconnector://{MYSQL_USER}:{MYSQL_PASSWORD}@{MYSQL_HOST}/{MYSQL_DB}"
    )
    # SWAGGER CONFIG
    API_TITLE = "Users API"
    API_VERSION = "v1"
    OPENAPI_VERSION = "3.0.3"
    OPENAPI_URL_PREFIX = "/"
    OPENAPI_SWAGGER_UI_PATH = "/swagger/index.html"
    OPENAPI_SWAGGER_UI_URL = "https://cdn.jsdelivr.net/npm/swagger-ui-dist/"
    OPENAPI_REDOC_PATH = "/redoc"
    OPENAPI_REDOC_UI_URL = (
        "https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone"
    )


class DevelopmentConfig(ConfigBase):
    DEBUG = True
    APP_PORT = configEnv('APP_PORT')
    APP_HOST = configEnv('APP_HOST')


config = {"development": DevelopmentConfig}
