from fastapi import APIRouter, Request


router = APIRouter()


@router.post("/create")
def create_user():
    return {"message": "User created"}
