from sqlalchemy import Column, Integer, String
from db.base import Base


class Comment(Base):
    """
        This Blog table responsible for holding analytics data for specific blog
    """
    __tablename__ = "comments"

    id = Column(Integer, primary_key=True)
    blog_id = Column(Integer)
    user_id = Column(Integer)
    content = Column(String)
    created_at = Column(Integer)
    updated_at = Column(Integer)
    deleted_at = Column(Integer)
