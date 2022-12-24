docker-compose down --volumes --remove-orphans && \
  docker-compose rm && \
  docker-compose build && \
  docker-compose up
# for f in $(find . -name go.mod)
# do (cd $(dirname $f); go mod tidy)
# done