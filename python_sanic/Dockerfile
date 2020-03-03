FROM python:3.7-alpine

RUN apk add --update libffi-dev openssl-dev musl-dev make g++ gcc

COPY requirements.txt /tmp/requirements.txt

RUN pip3 install -r /tmp/requirements.txt

WORKDIR /app

ADD run.py /app

CMD python3 run.py
