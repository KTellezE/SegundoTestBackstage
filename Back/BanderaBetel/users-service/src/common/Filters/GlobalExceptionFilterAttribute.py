from ..Dtos.ErrorMessageDto import ErrorMessageDto
from flask import jsonify

class GlobalExceptionFilterAttribute():
    # Manejador de errores para tu excepción personalizada
    def handle_custom_exception(error):
        # Verifica si la excepción tiene el atributo httpStatus
        if hasattr(error, 'httpStatus') and hasattr(error, 'message'):
            errorMessage = ErrorMessageDto(error.httpStatus, error.message, "", error.message).__dict__
        else:
            # Manejo genérico para excepciones que no tienen estos atributos
            errorMessage = ErrorMessageDto(500, "Error interno del servidor", "", str(error)).__dict__
        
        return jsonify(errorMessage), errorMessage['errorCode']