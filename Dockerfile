FROM ubuntu:xenial
WORKDIR /app
COPY japare .
COPY frontend/dist/ .
CMD ["/app/japare"]
