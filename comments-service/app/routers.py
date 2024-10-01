from fastapi import APIRouter
from fastapi.encoders import jsonable_encoder
from base import session_local
from models import Comment
from repository import CommentRepository
from services import CommentService
from schemas import Comment as CommentSchema


comment_router = APIRouter()
session = session_local()


@comment_router.get("/comments/", response_model=list[CommentSchema], status_code=200)
async def get_comments():
    repo = CommentRepository(session, Comment)
    comment = CommentService(comment_repo=repo)
    return comment.get_comments()


@comment_router.post("/comments/", response_model=CommentSchema, status_code=201)
async def add_comment(comment: CommentSchema):
    repo = CommentRepository(session, Comment)
    comment_service = CommentService(comment_repo=repo)
    comment_model = Comment(**comment.dict())
    comment_service.add_comment(comment_model)
    return comment


@comment_router.delete("/comments/{comment_id}", status_code=204)
async def delete_comment(comment_id: int):
    repo = CommentRepository(session, Comment)
    comment = CommentService(comment_repo=repo)
    comment.delete_comment(comment_id)


@comment_router.put("/comments/{comment_id}", status_code=200, response_model=CommentSchema)
async def update_comment(comment_id: int, comment_body: CommentSchema):
    repo = CommentRepository(session, Comment)
    comment = CommentService(comment_repo=repo)
    comment_json = jsonable_encoder(comment_body)
    comment.update_comment(comment_id, comment_json)
    return comment_body


@comment_router.get("/comments/{comment_id}", response_model=CommentSchema, status_code=200)
async def get_comment(comment_id: int):
    repo = CommentRepository(session, Comment)
    comment = CommentService(comment_repo=repo)
    comment = comment.get_comment(comment_id)
    return comment
