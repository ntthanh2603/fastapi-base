from fastapi import APIRouter
from src.api import users, home

router = APIRouter()

router.include_router(home.router, tags=["home"], prefix="/home")
router.include_router(users.router, tags=["users"], prefix="/users")
