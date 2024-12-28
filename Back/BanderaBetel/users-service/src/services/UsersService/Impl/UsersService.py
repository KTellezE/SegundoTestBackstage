from ..IUsersService import IUsersService
from ....persistence.DAO.Users import IUsersDao
from ....model.UsersModel import UsersModel
from ....common.Dtos.Response.UsersDto import UsersDto
from ....common.Util.ValidatorUtil import ValidatorUtil
from ....common.Dtos.Request.UsersRequestDto import UsersRequestDto
from datetime import datetime
from injector import inject

class UsersService(IUsersService):

    @inject
    def __init__(self, UsersDao: IUsersDao):
        self.UsersDao = UsersDao

    def GetAllAsync(self) -> list[UsersDto]:
        users = self.UsersDao.GetAllAsync()
        result = [UsersDto(user.Id, user.RoleId, user.Name, user.Username, user.Email,
                         user.UserCreated, user.Created, user.UserModified,
                         user.Modified, user.Active).__dict__
        for user in users]
        return result
    
    def GetByIdAsync(self, id: int) -> UsersDto:
        model = self.UsersDao.GetByIdAsync(id)
        ValidatorUtil.ValidateRequired(model, "El usuario no existe en la base de datos.")
            
        result = UsersDto(model.Id, model.RoleId, model.Name, model.Username, model.Email,
            model.UserCreated, model.Created, model.UserModified,
            model.Modified, model.Active).__dict__
        
        return result
    
    def InsertAsync(self, request: UsersRequestDto) -> bool:
        ValidatorUtil.ValidateNotRequired(
            self.UsersDao.GetByEmailAsync(request["email"]),
            "El correo electronico ya existe.")
        
        now = datetime.now()
        model = UsersModel(
            RoleId = request["roleId"],
            Name = request["name"],
            Username = request["username"],
            Email = request["email"]
        )
        
        model.Id = None
        model.Created = now
        model.UserCreated = "system"
        model.Modified = now
        model.UserModified = "system"
        model.Active = True
        
        return self.UsersDao.InsertAsync(model) 
    
    def UpdateAsync(self, id: int, request: UsersRequestDto) -> bool:
        dto = UsersDto(
            request["id"],
            request["roleId"],
            request["name"],
            request["username"],
            request["email"],
            "",
            datetime.now(),
            "",
            datetime.now(),
            request["active"])
        
        model = self.UsersDao.GetByIdAsync(id)
        ValidatorUtil.ValidateRequired(model, "El usuario no existe en la base de datos.")
        
        model.RoleId = dto.RoleId
        model.Name = dto.Name
        model.Username = dto.Username
        model.Email = dto.Email
        model.Modified = datetime.now()
        model.UserModified = "system"
        model.Active = dto.Active
        
        response = self.UsersDao.UpdateAsync()
        return response
    
    def DeleteAsync(self, id: int) -> bool:
        model = self.UsersDao.GetByIdAsync(id)
        ValidatorUtil.ValidateRequired(model, "El usuario no existe en la base de datos.")
        
        model.Modified = datetime.now()
        model.UserModified = "system"
        model.Active = False
        
        response = self.UsersDao.UpdateAsync()
        return response
    
