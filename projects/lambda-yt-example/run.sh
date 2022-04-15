docker build . -t work && \
docker run -v $(pwd):/go/src/lambda-yt-example -p 8080:8080 -it work bash