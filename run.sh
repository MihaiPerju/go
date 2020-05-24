docker build . -t work && \
docker run -v $(pwd):/work -it work bash