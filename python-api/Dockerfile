FROM ubuntu:20.04

WORKDIR /python-api

ENV FLASK_APP=server.py
ENV FLASK_RUN_HOST=0.0.0.0

COPY ./requirements.txt ./

RUN apt-get update && apt-get -y install python3.8 \ 
    python3-pip
RUN pip install --no-cache-dir -r requirements.txt


EXPOSE 5000

COPY . ./app/api

CMD ["flask","run"]
