FROM python:3.9

WORKDIR /app

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY ./app .

CMD ["uvicorn", "main:app", "--reload", "--port", "8002"]

# CMD ["u", "dev", "main.py", "--port", "8002"]
