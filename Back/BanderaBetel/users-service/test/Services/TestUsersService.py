from src.services.UsersService.Impl.UsersService import UsersService
from unittest import TestCase
from src.model.UsersModel import UsersModel
from config import config
from src.api import InitController
from ..BaseTest import BaseTest
from .Util.FunctionsUtil import FunctionsUtil
from src.common.Exceptions.BusinessException import BusinessException
from src.persistence.DAO.Users.Impl.UsersDao import UsersDao
from src.persistence.Db.database import db

class TestUsersService(TestCase, BaseTest):   
    
    @classmethod
    def setUpClass(self):
        configuration = config["development"]
        configuration.SQLALCHEMY_DATABASE_URI = 'sqlite:///:memory:'
        configuration.SQLALCHEMY_TRACK_MODIFICATIONS = False
        
        self.controller = InitController()
        self.app = self.controller.init_app(configuration)
        
        self.app_context = self.app.app_context()
        self.app_context.push()
        db.create_all()

        self.mockUsersDao = UsersDao()
        self.UsersService = UsersService(self.mockUsersDao)
        
    def setUp(self):
        FunctionsUtil.DeleteAllRecords()
        db.session.add_all(self.GetAllUsersModel())
        db.session.commit()
        
        
    def test_GetAllAsync(self):
        response = self.UsersService.GetAllAsync()
        self.assertIsNotNone(response)
        self.assertEqual(2, len(response))
     
    def test_GetByIdAsync(self):
        id = 1
        username = "user1"
        response = self.UsersService.GetByIdAsync(id)
        
        self.assertIsNotNone(response)
        self.assertEqual(id, response["Id"])
        self.assertEqual(username, response["Username"])
        
    def test_InsertAsync(self):
        data = self.GetUsersRequestDto()
        
        before = UsersModel.query.all()
        response = self.UsersService.InsertAsync(data)
        after = UsersModel.query.all()
        
        self.assertIsNotNone(response)
        self.assertTrue(response)
        self.assertEqual(len(before) + 1, len(after))
        
    def test_InsertAsyncBusinessException(self):
        data = self.GetUsersRequestDto()
        data["email"] = "user1@gmail.com"
        with self.assertRaises(BusinessException):
            self.UsersService.InsertAsync(data)
        
    def test_UpdateAsync(self):
        id = 1
        username = "user01"
        data = {
            "id": id,
            "email": "user1@gmail.com",
            "name": "usuario 1",
            "roleId": 1,
            "username": username,
            "active": True
        }
        
        response = self.UsersService.UpdateAsync(id, data)
        after = self.controller.db.session.get(UsersModel, id)
        
        self.assertIsNotNone(response)
        self.assertTrue(response)
        self.assertEqual(username, after.Username)
        
    def test_UpdateAsyncBusinessException(self):
        id = 100
        username = "user01"
        data = {
            "id": id,
            "email": "user1@gmail.com",
            "name": "usuario 1",
            "roleId": 1,
            "username": username,
            "active": True
        }
        
        with self.assertRaises(BusinessException):
            self.UsersService.UpdateAsync(id, data)
        
    def test_DeleteAsync(self):
        id = 2
        response = self.UsersService.DeleteAsync(id)
        after = db.session.get(UsersModel, id)
        
        self.assertIsNotNone(response)
        self.assertTrue(response)
        self.assertFalse(after.Active)
        
    def test_DeleteAsyncBusinessException(self):
        id = 100
        with self.assertRaises(BusinessException):
            self.UsersService.DeleteAsync(id)