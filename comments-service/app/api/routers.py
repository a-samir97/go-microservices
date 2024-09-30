from fastapi import APIRouter

comment_router = APIRouter()


@comment_router.get("/comments/")
async def get_comments():
    return {"message": "Get all comments"}


@comment_router.post("/comments/")
async def add_comment():
    return {"message": "Add a comment"}


@comment_router.delete("/comments/{comment_id}")
async def delete_comment(comment_id: int):
    return {"message": "Delete a comment"}


@comment_router.put("/comments/{comment_id}")
async def update_comment(comment_id: int):
    return {"message": "Update a comment"}


@comment_router.get("/comments/{comment_id}")
async def get_comment(comment_id: int):
    return {"message": "Get a comment"}
