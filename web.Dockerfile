FROM debian:latest

RUN apt-get update && apt-get install -y tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

WORKDIR /app

COPY web /app
COPY front/  /app/front
COPY data/  /app/data

COPY entrypoint.sh  /app

RUN chmod +x /app/entrypoint.sh

CMD ["./entrypoint.sh"]