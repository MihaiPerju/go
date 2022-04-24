docker build . -t work && \
docker run -v $(pwd):/go/src/wolfram-witai -p 8080:8080 -it work bash