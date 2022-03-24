docker build . -t work && \
docker run -v $(pwd):/go/src/go-movies-crud -p 8000:8000 -it work bash