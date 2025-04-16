from sqlalchemy import Column, String, DateTime, Integer, Enum
from sqlalchemy.ext.declarative import declarative_base
import datetime
import uuid
from helpers.role import RoleType
from models.base import BareBaseModel


class User(BareBaseModel):
    id = Column(String, primary_key=True, default=lambda: str(uuid.uuid4()))
    email = Column(String, nullable=False, unique=True)
    password = Column(String, nullable=False)
    username = Column(String, nullable=False, unique=True)
    bio = Column(String, nullable=True)
    role = Column(Enum(RoleType), nullable=False, default=RoleType.USER)
    refresh_token = Column(String, nullable=True)
