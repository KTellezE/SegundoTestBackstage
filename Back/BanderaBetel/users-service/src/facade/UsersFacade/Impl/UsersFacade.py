from ..IUsersFacade import IUsersFacade
from ....services.UsersService.Impl.UsersService import UsersService
from ....services.UsersService.Impl.UsersService import IUsersService
from ....common.Dtos.Response.UsersDto import UsersDto
from ....common.Dtos.Request.UsersRequestDto import UsersRequestDto
from injector import inject

class UsersFacade(IUsersFacade):

    UsersFacade: IUsersFacade
    UsersService: UsersService

    @inject
    def __init__(self, UsersService: IUsersService):
        self.UsersService = UsersService

    def GetAllAsync(self) -> list[UsersDto]:
        return self.UsersService.GetAllAsync()

    def GetByIdAsync(self, id: int) -> UsersDto:
        return self.UsersService.GetByIdAsync(id)

    def InsertAsync(self, request: UsersRequestDto) -> bool:
        return self.UsersService.InsertAsync(request)

    def UpdateAsync(self, id: int, request: UsersRequestDto) -> bool:
        return self.UsersService.UpdateAsync(id, request)

    def DeleteAsync(self, id: int) -> bool:
        return self.UsersService.DeleteAsync(id)
