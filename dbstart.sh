#!/bin/bash
PROJECT=courses
MYSQL_PASSWORD=test
DATABASENAME=courses
PHP_ADMIN=myadmin
USER_NAME=root
PASSWORD=test

docker stop ${PHP_ADMIN}
docker rm ${PHP_ADMIN}
docker stop ${PROJECT}
docker rm ${PROJECT}
docker pull mysql
docker pull phpmyadmin/phpmyadmin

docker run -d --name=${PROJECT} \
    --restart always \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=${PASSWORD} \
    mysql

docker run --name myadmin --restart always -d --link ${PROJECT}:db -p 8080:80 phpmyadmin/phpmyadmin

# wait for DB up
sleep 15
docker exec -it ${PROJECT} mysql -u${USER_NAME} -hlocalhost -P 3306 -p -e "create database ${DATABASENAME};"
