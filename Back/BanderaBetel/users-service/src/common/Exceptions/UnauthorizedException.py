from http import HTTPStatus

class UnauthorizedException(Exception):
    def __init__(self, message, httpStatus = HTTPStatus.UNAUTHORIZED):
        self.message = message
        self.httpStatus = httpStatus

        super().__init__(self.message)