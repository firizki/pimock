FROM alpine:3.10.3

RUN wget https://github.com/firizki/ailea/releases/download/v0.0.1/linux-amd64.zip
RUN unzip linux-amd64.zip -d /usr/local/bin/
ADD responses/ responses/

EXPOSE 8080

ENTRYPOINT ["ailea"]
