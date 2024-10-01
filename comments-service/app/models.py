from sqlalchemy import Column, Integer, String, DateTime
from base import Base


class Comment(Base):
    """
        This Comment table responsible for holding comments for specific blog
    """
    __tablename__ = "comments"

    id = Column(Integer, primary_key=True)
    blog_id = Column(Integer)
    user_id = Column(Integer)
    content = Column(String)
    created_at = Column(DateTime(timezone=True), server_default="now()")
    updated_at = Column(DateTime(timezone=True), onupdate="now()")
    deleted_at = Column(DateTime)
