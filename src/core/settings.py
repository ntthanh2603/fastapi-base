from pydantic_settings import BaseSettings


class Settings(BaseSettings):
  # Database Postgres
  DATABASE_HOST: str
  DATABASE_PORT: int
  DATABASE_USERNAME: str
  DATABASE_PASSWORD: str
  DATABASE_NAME: str

  # Database Redis
  REDIS_PASSWORD: str
  REDIS_DB: int
  REDIS_HOST: str
  REDIS_PORT: int

  # JWT
  JWT_ALGORITHM: str
  JWT_SECRET_KEY: str
  JWT_ACCESS_EXPIRE: int
  JWT_REFRESH_EXPIRE: int
  JWT_REFRESH_EXPIRE_DAY: int


  # FastAPI
  PORT: int
  HOST: str
  API_PREFIX: str

  class Config:
    env_file = ".env"


settings = Settings()