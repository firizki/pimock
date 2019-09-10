FROM ubuntu:18.04

COPY ailea /usr/local/bin/ailea
ADD responses/ responses/

EXPOSE 8080

ENTRYPOINT ["ailea"]
