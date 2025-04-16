from fastapi import APIRouter, Request
from src.services.auth import AuthService
from src.schemas.user import CreateUserDto
from src.services.users import UserService

router = APIRouter()
auth_service = AuthService()
user_service = UserService()

@router.get("/")
def create_user(req: Request):
    return {"message": "User created",
            "user": req.user}

@router.post("/create")
async def create_user(dto: CreateUserDto):
    user = await user_service.create_user(dto)
    return user