from unittest.mock import MagicMock
from src.facade.UsersFacade.Impl.UsersFacade import UsersFacade
from src.common.Dtos.Response.UsersDto import UsersDto
from unittest import TestCase
from ..BaseTest import BaseTest
from src.services.UsersService.IUsersService import IUsersService

class TestUsersFacade(TestCase, BaseTest):
    
    def setUp(self):
        self.mockUsersService = MagicMock(return_value = IUsersService)
        self.UsersFacade = UsersFacade(self.mockUsersService)
        
        data = UsersDto(1, 1, "Usuario 1", "user1", "user1@gmail.com", "system", "2023-08-03 14:40:10", "system", "2023-08-03 14:40:10", True)
        list = [data,]
        
        self.UsersFacade.UsersService.GetAllAsync = MagicMock(return_value = list)
        self.UsersFacade.UsersService.GetByIdAsync = MagicMock(return_value = data)
        self.UsersFacade.UsersService.InsertAsync = MagicMock(return_value = True)
        self.UsersFacade.UsersService.UpdateAsync = MagicMock(return_value = True)
        self.UsersFacade.UsersService.DeleteAsync = MagicMock(return_value = True)
    
    def test_GetAllAsync(self):
        response = self.UsersFacade.GetAllAsync()
        
        self.assertIsNotNone(response)
        self.assertEqual(1, len(response))
        
    def test_GetByIdAsync(self):
        id = 1
        response = self.UsersFacade.GetByIdAsync(id)
        
        self.assertIsNotNone(response)
        self.assertEqual(id, response.Id)
        
    def test_InsertAsync(self):
        data = self.GetUsersRequestDto()
        response = self.UsersFacade.InsertAsync(data)
        
        self.assertIsNotNone(response)
        self.assertTrue(response)
        
    def test_UpdateAsync(self):
        id = 1
        data = self.GetUsersRequestDto()
        data["id"] = id
        response = self.UsersFacade.UpdateAsync(id, data)
        
        self.assertIsNotNone(response)
        self.assertTrue(response)
        
    def test_DeleteAsync(self):
        id = 1
        response = self.UsersFacade.DeleteAsync(id)
        
        self.assertIsNotNone(response)
        self.assertTrue(response)