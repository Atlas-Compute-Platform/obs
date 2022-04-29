# Atlas Dockerfile
FROM alpine:latest
RUN apk add go
COPY ./obs /usr/local/bin/obs
EXPOSE 8800/tcp
CMD ["/usr/local/bin/obs"]
