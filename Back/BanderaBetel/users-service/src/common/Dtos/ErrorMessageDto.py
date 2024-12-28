class ErrorMessageDto(): 
    def __init__(self, errorCode, errorMessage, userError, info):
        self.errorCode = errorCode
        self.errorMessage = errorMessage
        self.userError = userError
        self.info = info