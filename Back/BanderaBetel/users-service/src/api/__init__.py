from flask import Flask
from flask_smorest import Api # type: ignore
from flask_sqlalchemy import SQLAlchemy # type: ignore
from .DependencyInjection import injector

class InitController():
    
    def __init__(self):
        self.db:SQLAlchemy = injector.get(SQLAlchemy)
        self.app:Flask = injector.get(Flask)
        
    def init_app(self, config) -> Flask:
        #App config
        self.app.config.from_object(config)
        
        #Db config
        self.db.init_app(self.app)
        
        #Swagger blueprint
        from .controllers import router as routerMapped
        api:Api = Api(self.app)
        api.register_blueprint(routerMapped)
        
        return self.app
    