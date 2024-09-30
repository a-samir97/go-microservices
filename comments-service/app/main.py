import asyncio
import logging
from api.routers import comment_router
from fastapi import FastAPI

# from messaging.consumer import start_consumer

app = FastAPI()

loop = asyncio.get_event_loop()
logging.basicConfig(level=logging.INFO)

app.include_router(comment_router, prefix="/api/v1")


@app.get("/")
def read_root():
    return {"message": "Comments Service"}


# @app.on_event("startup")
# async def startup_event():
#     loop.create_task(start_consumer(loop))
