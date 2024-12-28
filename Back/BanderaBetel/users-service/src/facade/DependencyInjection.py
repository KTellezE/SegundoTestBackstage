from injector import Binder, Module, singleton
from .UsersFacade.IUsersFacade import IUsersFacade
from .UsersFacade.Impl.UsersFacade import UsersFacade
from ..persistence.DAO.Users.Impl.UsersDao import UsersDao
from ..services.UsersService.Impl.UsersService import UsersService
from ..facade.UsersFacade.IUsersFacade import IUsersFacade
from ..facade.UsersFacade.Impl.UsersFacade import UsersFacade

class AddFacade(Module):
    
    def configure(self, binder: Binder):
        binder.bind(IUsersFacade, to=UsersFacade(UsersService(UsersDao())), scope=singleton)
        
        