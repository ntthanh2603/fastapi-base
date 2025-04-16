from fastapi import FastAPI
from starlette.middleware.cors import CORSMiddleware
from src.api.router import router
import os
import uvicorn
import asyncio
from src.middlewares.auth import AuthMiddleware
from src.core.settings import settings



def get_application() -> FastAPI:
    # Create the FastAPI application
    application = FastAPI(
        title="Social network SNet",
        docs_url="/docs",
        redoc_url="/re-docs",
        openapi_url=f"{settings.API_PREFIX}/openapi.json",
        description="""
        API use FastAPI docs for Social network SNet:
        - Posts.
        - Recomendations.
        - Search vector.
        - Suggest user follow.
        """,
    )

    application.middleware("http")(AuthMiddleware)

    application.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )
    application.include_router(router, prefix=settings.API_PREFIX)
    return application


app = get_application()


def main():

    config = uvicorn.Config(
        app=app,
        host=settings.HOST,
        port=int(settings.PORT_FASTAPI),
        reload=True,
    )

    server = uvicorn.Server(config)
    asyncio.run(server.serve())


if __name__ == "__main__":
    main()
