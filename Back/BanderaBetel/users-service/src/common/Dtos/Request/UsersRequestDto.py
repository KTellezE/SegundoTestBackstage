from marshmallow import Schema, fields

class UsersRequestDto(Schema): 
    id = fields.Int(required=False)
    email = fields.Str(required=True)
    name = fields.Str(required=True)
    roleId = fields.Int(required=True)
    username = fields.Str(required=True)
    active = fields.Boolean(required=True)