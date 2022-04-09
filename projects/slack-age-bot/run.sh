docker build . -t work && \
docker run -v $(pwd):/go/src/slack-age-bot -p 8000:8000 -it work bash