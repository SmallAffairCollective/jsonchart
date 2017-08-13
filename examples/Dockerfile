FROM python:2.7-slim
MAINTAINER defermat <defermat@defermat.net>

COPY *.json /examples/
WORKDIR /examples

EXPOSE 8000
ENTRYPOINT ["python", "-m", "SimpleHTTPServer"]
