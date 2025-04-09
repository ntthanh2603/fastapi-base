from pydantic import BaseModel
from typing import Optional
from fastapi import Request, HTTPException, status
from fastapi.responses import JSONResponse
import jwt


# Middleware authorization
async def AuthMiddleware(request: Request, call_next):
    # List routers public
    public_routes = {"/docs", "/openapi.json", "/token", "/home"}

    if request.url.path in public_routes:
        return await call_next(request)

    try:
        # Get token from header
        token = request.headers.get("authorization")
        if token:
            token = token.split(" ")[1]

        if not token:
            return JSONResponse(
                status_code=status.HTTP_401_UNAUTHORIZED,
                content={"message": "Authorization header missing or invalid"},
            )

        return call_next(request)

    except HTTPException as e:
        return JSONResponse(status_code=e.status_code, content={"detail": e.detail})
