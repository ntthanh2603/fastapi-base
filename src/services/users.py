import uuid
from src.schemas.user import CreateUserDto
from src.db.postgres import get_db
from sqlalchemy import text
from fastapi import HTTPException, status
import bcrypt


class UserService:

    # Hash password
    def hash_password(self, password) -> str:
        salt = bcrypt.gensalt()       
        return bcrypt.hashpw(password.encode('utf-8'), salt).decode('utf-8')
    
    # Verify password
    def verify_password(self, password, hashed_password) -> bool:
        return bcrypt.checkpw(password.encode('utf-8'), hashed_password.encode('utf-8'))
    
    # Create user
    async def create_user(self, dto: CreateUserDto):
        async for db in get_db():
            try:
                # Check if user already exists
                query = text("SELECT * FROM users WHERE email = :email OR username = :username")
                result = await db.execute(query, {"email": dto.email, "username": dto.username})
                user = result.mappings().first()
                print(f'User: {user}')
                if user:
                    raise HTTPException(
                        status_code=status.HTTP_400_BAD_REQUEST,
                        detail="User already exists",
                    )

                query = text("""
                    INSERT INTO users (id, email, password, username, bio)
                    VALUES (:id, :email, :password, :username, :bio)
                    RETURNING *
                """)

                # Hash password
                password = self.hash_password(dto.password)

                params = {
                    "id": str(uuid.uuid4()),
                    "email": dto.email,
                    "password": password,
                    "username": dto.username,
                    "bio": dto.bio,
                }
                # Save user
                result = await db.execute(query, params)
                await db.commit()
                user = result.mappings().first()

                return user

            except HTTPException as http_exc:
                raise http_exc
            except Exception as e:
                raise HTTPException(
                    status_code=500,
                    detail=f"Database error: {str(e)}"
                )
