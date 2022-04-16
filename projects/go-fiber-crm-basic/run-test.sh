docker build . -t work && \
docker run  -p 8000:8000 -it work bash
# -v $(pwd):/go/src/go-fiber-crm-basic