from aiokafka import AIOKafkaConsumer
import logging

logging.basicConfig(level=logging.INFO)


async def start_consumer(loop):
    consumer = AIOKafkaConsumer(
        'blog_viewed',
        bootstrap_servers='kafka:9092',
        auto_offset_reset='earliest',
        enable_auto_commit=True,
        loop=loop
    )

    try:
        await consumer.start()
        logging.info("Consumer started")
        print("Consumer started", '\n\n')
        async for message in consumer:
            logging.info("Message received: %s", message.value.decode())

    finally:
        await consumer.stop()
        logging.info("Consumer stopped")
