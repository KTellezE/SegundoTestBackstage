from abc import ABC, abstractmethod
from ....model.UsersModel import UsersModel

class IUsersDao(ABC):

    @abstractmethod
    def GetAllAsync(self) -> list[UsersModel]:
        pass

    @abstractmethod
    def GetByIdAsync(self, id: int) -> UsersModel:
        pass

    @abstractmethod
    def GetByEmailAsync(self, email: str) -> UsersModel:
        pass

    @abstractmethod
    def InsertAsync(self, model: UsersModel) -> bool:
        pass

    @abstractmethod
    def UpdateAsync(self) -> bool:
        pass
