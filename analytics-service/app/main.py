import asyncio
import logging
from typing import Union

from fastapi import FastAPI
from pydantic import BaseModel

from messaging.consumer import start_consumer

app = FastAPI()

loop = asyncio.get_event_loop()
logging.basicConfig(level=logging.INFO)


class Item(BaseModel):
    name: str
    price: float
    is_offer: Union[bool, None] = None


@app.get("/")
def read_root():
    return {"hello": "World"}


@app.get("/items/{item_id}/")
def read_item(item_id: int, q: Union[str, None] = None):
    return {"item_id": item_id, "q": q}


@app.put("/items/{item_id}/")
def update_item(item_id: int, item: Item):
    return {"item_name": item.name, "item_id": item_id}


@app.on_event("startup")
async def startup_event():
    loop.create_task(start_consumer(loop))
