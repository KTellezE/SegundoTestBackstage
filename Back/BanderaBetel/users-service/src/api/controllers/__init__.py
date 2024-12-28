from flask_smorest import Blueprint
from  ...common.Filters.GlobalExceptionFilterAttribute import GlobalExceptionFilterAttribute

router = Blueprint(
    "Users", "UsersController", url_prefix="/api/Users", description="Users API"
)

router.register_error_handler(Exception, GlobalExceptionFilterAttribute.handle_custom_exception)
from src.api.controllers.UsersController import *