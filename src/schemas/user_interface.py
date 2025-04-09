from pydantic import BaseModel
from typing import Optional


class IUser(BaseModel):
    id: str
    role: str
    iat: int
    exp: int
