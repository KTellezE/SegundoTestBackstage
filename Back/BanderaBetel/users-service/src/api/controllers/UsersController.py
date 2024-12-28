from flask.views import MethodView
from ...facade.UsersFacade.Impl.UsersFacade import IUsersFacade
from ...common.Dtos.Request.UsersRequestDto import UsersRequestDto
from flask import jsonify
from . import router
from ..DependencyInjection import injector

class UsersController(MethodView):

    UsersFacade: IUsersFacade

    @router.route("", methods=["GET"])
    @router.response(status_code=200)
    def GetAllAsync():
        UsersFacade = injector.get(IUsersFacade)
        return jsonify(UsersFacade.GetAllAsync())
    
    @router.route("/<int:id>", methods=["GET"])
    @router.response(status_code=200)
    def GetByIdAsync(id: int):
        UsersFacade = injector.get(IUsersFacade)
        return jsonify(UsersFacade.GetByIdAsync(id))
    
    @router.route("", methods=["POST"])
    @router.response(201, UsersRequestDto)
    @router.arguments(UsersRequestDto)
    def InsertAsync(request: UsersRequestDto):
        UsersFacade =  injector.get(IUsersFacade)
        return jsonify(UsersFacade.InsertAsync(request))

    @router.route("/<int:id>", methods=["PUT"])
    @router.response(201, UsersRequestDto)
    @router.arguments(UsersRequestDto)
    def UpdateAsync(request: UsersRequestDto, id: int):
        UsersFacade =  injector.get(IUsersFacade)
        return jsonify(UsersFacade.UpdateAsync(id, request))
    
    @router.route("/<int:id>", methods=["DELETE"])
    @router.response(status_code=201)
    def DeleteAsync(id: int):
        UsersFacade =  injector.get(IUsersFacade)
        return jsonify(UsersFacade.DeleteAsync(id))
    
    @router.route("/ping", methods=["GET"])
    @router.response(status_code=200)
    def Ping():
        return jsonify("pong")
