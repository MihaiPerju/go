docker build . -t work && \
docker run -v $(pwd):/go/src/email-checker-tool -p 8000:8000 -it work bash