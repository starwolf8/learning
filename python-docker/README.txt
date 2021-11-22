Reference: https://docs.docker.com/language/python/run-containers/

Create a virtual env: `virtualenv <NAME_OF_VIRTENV>`

Activate Env: `source bin/activate`

Exit/Deactivate Env: `deactivate`

Install Flask while in Env: `pip install flask`

====Prep for Dockerfile====

view images: `docker images`

build a python image: `docker build --tag python-docker .`

run docker image : `docker run python-docker`

run docker image with port access: `docker run --publish 5001:5000 python-docker`