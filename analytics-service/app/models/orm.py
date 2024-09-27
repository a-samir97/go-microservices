from sqlalchemy import Column, Integer
from db.base import Base


class Blog(Base):
    """
        This Blog table responsible for holding analytics data for specific blog
    """
    __tablename__ = "blogs"

    id = Column(Integer, primary_key=True)
    blog_id = Column(Integer)
    viewed = Column(Integer)
    likes = Column(Integer)
    dislikes = Column(Integer)
