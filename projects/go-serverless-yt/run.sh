docker build . -t work && \
docker run -v $(pwd):/go/src/go-serverless-yt -p 8080:8080 -it work bash