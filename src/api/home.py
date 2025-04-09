from fastapi import APIRouter

router = APIRouter()


@router.get("")
def create_user():
    return {"message": "Hello, World! 🌍"}
