from http import HTTPStatus

class NotFoundException(Exception):
    def __init__(self, message, httpStatus = HTTPStatus.NOT_FOUND):
        self.message = message
        self.httpStatus = httpStatus

        super().__init__(self.message)