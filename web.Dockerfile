FROM debian:latest

WORKDIR /app

COPY web /app
COPY front/  /app/front
COPY data/  /app/data

COPY entrypoint.sh  /app

RUN chmod +x /app/entrypoint.sh

CMD ["./entrypoint.sh"]