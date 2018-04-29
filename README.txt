#Use the following to build as a docker image (working dir is the project root)
docker build -t golang-test src/mongotest

# Use the following to run as a docker container
docker run -p 8000:8000 -h golang-server golang-test