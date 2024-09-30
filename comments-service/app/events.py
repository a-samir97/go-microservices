import asyncio

from fastapi import FastAPI
from messaging.consumer import start_consumer


def create_startup_event(app: FastAPI):
    @app.on_event("startup")
    async def startup_event():
        asyncio.timeout(10)
        loop = asyncio.get_event_loop()
        loop.create_task(start_consumer())
        print("Consumer started")


def create_shutdown_event(app: FastAPI):
    @app.on_event("shutdown")
    def shutdown_event():
        print("Consumer stopped")
