FROM alpine:latest
WORKDIR /app
ADD ./dist /app
RUN chmod +x ./server-green
EXPOSE 4000
ENTRYPOINT [ "/app/server-green" ]