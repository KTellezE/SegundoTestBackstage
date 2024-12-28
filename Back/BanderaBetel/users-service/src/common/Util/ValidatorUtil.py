from ..Exceptions.BusinessException import BusinessException

class ValidatorUtil():
    
    def ValidateRequired(value, message):
        if value is None:
            raise BusinessException(message)
        
    def ValidateNotRequired(value, message):
        if not value is None:
            raise BusinessException(message)
        
    def ValidateCondition(condition, message):
        if condition is True:
            raise BusinessException(message)