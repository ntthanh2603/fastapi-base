from pydantic import BaseModel
from typing import Optional

class UserBase(BaseModel):
    username: str

class IUser(UserBase):
    id: str
    role: str
    iat: int
    exp: int

class CreateUserDto(UserBase):
    email: str
    password: str
    bio: Optional[str] = None
