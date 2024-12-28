from ..persistence.Db.database import db

class UsersModel(db.Model):
    __tablename__ = 'tbl_user'
    
    Id = db.Column('cd_id', db.Integer, primary_key=True, autoincrement=True)
    RoleId = db.Column('cd_role', db.Integer, nullable=False)
    Name = db.Column('nb_name', db.String(100), nullable=False)
    Username = db.Column('nb_username', db.String(100), nullable=False)
    Email = db.Column('nb_email', db.String(100), nullable=False)
    UserCreated = db.Column('nb_usercreated', db.String(100), nullable=False)
    Created = db.Column('ts_created', db.DateTime, nullable=False)
    UserModified = db.Column('nb_usermodified', db.String(100), nullable=False)
    Modified = db.Column('ts_modified', db.DateTime, nullable=False)
    Active = db.Column('st_active', db.Boolean, nullable=False)
