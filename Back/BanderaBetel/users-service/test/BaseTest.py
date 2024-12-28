from abc import ABC
from src.common.Dtos.Request.UsersRequestDto import UsersRequestDto
from src.common.Dtos.Response.UsersDto import UsersDto
from src.model.UsersModel import UsersModel
from datetime import datetime

class BaseTest(ABC):
    def GetUsersRequestDto(self) -> UsersRequestDto:
        return {
            "email": "user3@gmail.com",
            "name": "usuario 3",
            "roleId": 1,
            "username": "user3",
            "active": True
        }
        
    def GetUsersDto(self) -> UsersDto:
        return UsersDto(1, 1, "Usuario 1", "user1", "user1@gmail.com", "system", "2023-08-03 14:40:10", "system", "2023-08-03 14:40:10", True)
    
    def GetAllUsersModel(self) -> list[UsersModel]:
        return [
            UsersModel(Id = 1, RoleId = 1, Name = "usuario 1", Username = "user1", Email = "user1@gmail.com", UserCreated = "system", Created = datetime.now(), UserModified = "system", Modified = datetime.now(), Active = True ),
            UsersModel(Id = 2, RoleId = 1, Name = "usuario 2", Username = "user2", Email = "user2@gmail.com", UserCreated = "system", Created = datetime.now(), UserModified = "system", Modified = datetime.now(), Active = True )
        ]