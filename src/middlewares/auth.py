from fastapi import Request, HTTPException, status
from fastapi.responses import JSONResponse
from src.services.auth import AuthService

auth_service = AuthService()

# List routers public
PUBLIC_ROUTES = {"/docs", "/openapi.json", "/token", "/home", "/users/create"}



# Middleware authorization
async def AuthMiddleware(request: Request, call_next):

    if request.url.path in PUBLIC_ROUTES:
        return await call_next(request)

    # Get token from header
    token = request.headers.get("authorization")
    if token:
        token = token.split(" ")[1]

    if not token:
        return JSONResponse(
            status_code=status.HTTP_401_UNAUTHORIZED,
            content={"message": "Authorization header missing or invalid"},
            )
        
    try:
        # Verify token
        user = await auth_service.verify_token(token)

        request.user = user

        print(f'User: {user}')
        return call_next(request)

    except HTTPException as e:
        return JSONResponse(status_code=e.status_code, content={"detail": e.detail})
