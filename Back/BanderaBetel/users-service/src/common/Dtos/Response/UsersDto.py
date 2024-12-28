class UsersDto():    
    def __init__(self, id, roleId, name, username, email, userCreated, created, userModified, modified, active):
        self.Id = id
        self.RoleId = roleId
        self.Name = name
        self.Username = username
        self.Email = email
        self.UserCreated = userCreated
        self.Created = created
        self.UserModified = userModified
        self.Modified = modified
        self.Active = active