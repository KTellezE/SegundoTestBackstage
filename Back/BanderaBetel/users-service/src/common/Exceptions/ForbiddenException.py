from http import HTTPStatus

class ForbiddenException(Exception):
    def __init__(self, message, httpStatus = HTTPStatus.FORBIDDEN):
        self.message = message
        self.httpStatus = httpStatus

        super().__init__(self.message)