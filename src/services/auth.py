import jwt
from fastapi import  HTTPException, status
from src.schemas.user import IUser
from src.core.settings import settings



class AuthService:
    def __init__(self):
        pass
    # Decode token
    def decode_token(self, token: str) -> IUser:
        try:
            payload = jwt.decode(token, options={"verify_signature": False})

            return payload
        except jwt.PyJWTError:
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail="Invalid token",
                headers={"WWW-Authenticate": "Bearer"},
            )

    # Verify token
    async def verify_token(self, token: str) -> IUser:
        try:
            result = jwt.decode(
                token, settings.JWT_SECRET_KEY, algorithms=settings.JWT_ALGORITHM
            )

            return result

        except jwt.PyJWTError:
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail="Invalid token",
                headers={"WWW-Authenticate": "Bearer"},
            )
