FROM python:3.9

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

RUN pip install --upgrade pip

WORKDIR /app

COPY ./requirements.txt /app

RUN pip install --no-cache-dir -r requirements.txt

COPY . /app

CMD ["python", "./application.py"]
