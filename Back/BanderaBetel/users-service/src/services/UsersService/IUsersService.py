from abc import ABC, abstractmethod
from ...common.Dtos.Request.UsersRequestDto import UsersRequestDto
from ...common.Dtos.Response.UsersDto import UsersDto

class IUsersService(ABC):
    
    @abstractmethod
    def GetAllAsync(self) -> list[UsersDto]:
        pass

    @abstractmethod
    def GetByIdAsync(self, id: int) -> UsersDto:
        pass

    @abstractmethod
    def InsertAsync(self, request: UsersRequestDto) -> bool:
        pass
    
    @abstractmethod
    def UpdateAsync(self, id: int, request: UsersRequestDto) -> bool:
        pass

    @abstractmethod
    def DeleteAsync(self, id: int) -> bool:
        pass
