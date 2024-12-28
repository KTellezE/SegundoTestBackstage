from injector import Binder, Module, singleton
from .UsersService.IUsersService import IUsersService
from .UsersService.Impl.UsersService import UsersService
from ..persistence.DAO.Users.Impl.UsersDao import UsersDao

class AddServices(Module):
    
    def configure(self, binder: Binder):
        binder.bind(IUsersService, to=UsersService(UsersDao()), scope=singleton)
        