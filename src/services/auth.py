import jwt
from fastapi import Request, HTTPException, status
from fastapi.responses import JSONResponse
import os
from schemas.user import IUser


class AuthService:
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
            payload = self.decode_token(token)

            result = jwt.decode(
                token, secret_key, algorithms=[os.getenv("JWT_ALGORITHM")]
            )

            return result

        except jwt.PyJWTError:
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail="Invalid token",
                headers={"WWW-Authenticate": "Bearer"},
            )
