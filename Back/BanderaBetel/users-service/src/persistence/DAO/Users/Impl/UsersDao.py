from .....model.UsersModel  import UsersModel
from ..IUsersDao import IUsersDao
from .....model.UsersModel import db

class UsersDao(IUsersDao):
    
    def GetAllAsync(self) -> list[UsersModel]:
        return db.session.query(UsersModel).filter(UsersModel.Active == True).all()

    def GetByIdAsync(self, id: int) -> UsersModel:
        return db.session.get(UsersModel, id)

    def GetByEmailAsync(self, email: str) -> UsersModel:
        return db.session.query(UsersModel).filter(UsersModel.Email == email).first()

    def InsertAsync(self, model: UsersModel) -> bool:
        db.session.add(model)
        db.session.commit()
        return True

    def UpdateAsync(self) -> bool:
        db.session.commit()
        return True
