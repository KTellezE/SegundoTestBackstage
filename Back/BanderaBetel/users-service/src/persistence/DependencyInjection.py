from flask_sqlalchemy import SQLAlchemy # type: ignore
from injector import Binder, Module, singleton
from flask import Flask
from .DAO.Users.IUsersDao import IUsersDao
from .DAO.Users.Impl.UsersDao import UsersDao
from .Db.database import db
from ..model.UsersModel import UsersModel

class AddPersistence(Module):
    
    def configure(self, binder: Binder) -> None:
        binder.bind(SQLAlchemy, to=db, scope=singleton)
        binder.bind(Flask, to=Flask(__name__), scope=singleton)
        binder.bind(UsersModel, to=UsersModel(), scope=singleton)
        binder.bind(IUsersDao, to=UsersDao(), scope=singleton)
        
        