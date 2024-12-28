from injector import Module, Injector
from ..persistence.DependencyInjection import AddPersistence
from ..services.DependencyInjection import AddServices
from ..facade.DependencyInjection import AddFacade

class AddConfiguration(Module):
    
    def configure(self, binder):
        binder.install(AddPersistence())
        binder.install(AddServices())
        binder.install(AddFacade())
        
injector = Injector([AddConfiguration()])
        