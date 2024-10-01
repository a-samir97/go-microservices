from datetime import datetime
from pydantic import BaseModel


class Comment(BaseModel):
    blog_id: int
    user_id: int
    content: str
    created_at: datetime
    updated_at: datetime
    deleted_at: datetime

    class Config:
        orm_mode = True


class CommentResponse(BaseModel):
    status: str
    message: str
    data: Comment
    error: str

    class Config:
        orm_mode = True
