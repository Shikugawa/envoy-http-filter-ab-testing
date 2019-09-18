FROM alpine:latest
WORKDIR /app
ADD ./dist /app
RUN chmod +x ./server-blue
EXPOSE 8000
ENTRYPOINT [ "/app/server-blue" ]