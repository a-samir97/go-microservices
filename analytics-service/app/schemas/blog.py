from typing import Union
from pydantic import BaseModel


class BlogBase(BaseModel):
    id: int
    blog_id: int


class BlogViewed(BlogBase):
    viewed: int

    class Config:
        orm_mode = True


class BlogLikes(BlogBase):
    likes: int

    class Config:
        orm_mode = True


class BlogDislikes(BlogBase):
    dislikes: int

    class Config:
        orm_mode = True


class BlogClaps(BlogBase):
    claps: int

    class Config:
        orm_mode = True
