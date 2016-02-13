#! /bin/bash
# run a local docker container based on the image mysql:latest
docker run -d -v `pwd`/exp.sql:/home/exp.sql -v `pwd`/importSql.sh/:/home/importSql.sh --restart=always -e MYSQL_ALLOW_EMPTY_PASSWORD=true mysql:latest
