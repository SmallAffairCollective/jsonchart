FROM alpine:3.6
MAINTAINER defermat <defermat@defermet.net>

RUN apk add --update \
    python \
    py2-pip \
    && rm -rf /var/cache/apk/*

COPY . /generator
WORKDIR /generator
RUN pip install -r requirements.txt

EXPOSE 80

ENTRYPOINT ["python", "gen.py"]
CMD [""]
