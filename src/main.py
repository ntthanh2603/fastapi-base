from fastapi import FastAPI, Request, Depends
from starlette.middleware.cors import CORSMiddleware
from middlewares.auth_middleware import auth_middleware
import os
from api.api_router import router
import uvicorn
from database.base import get_db
import asyncio

def create_application() -> FastAPI:
    # Create the FastAPI application
    application = FastAPI(
        title='Social network SNet', docs_url="/docs", redoc_url='/re-docs',
        openapi_url=f"{os.getenv('API_PREFIX', '')}/openapi.json",
        description='''
        API use FastAPI docs for Social network SNet:
        - Posts.
        - Recomendations.
        - Search vector.
        - Suggest user follow.
        '''
    )
    
    application.middleware("http")(auth_middleware)

    application.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],  
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )
    application.include_router(router, prefix=os.getenv('API_PREFIX', ''))
    return application

app = create_application()

def main():
    config = uvicorn.Config(
        app=app,
        host=os.getenv('HOST', '0.0.0.0'),
        port=int(os.getenv('PORT', '8000')),
        reload=True
    )
  
    server = uvicorn.Server(config)
    asyncio.run(server.serve())

if __name__ == '__main__':
    main()
