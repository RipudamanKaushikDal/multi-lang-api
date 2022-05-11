FROM ubuntu:20.04

WORKDIR /app

COPY requirements.txt ./

RUN apt-get update && apt-get -y install python3.8 \ 
    python3-pip
RUN pip install --no-cache-dir -r requirements.txt

COPY . ./

EXPOSE 5000

