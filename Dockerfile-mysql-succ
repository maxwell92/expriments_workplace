# This is Dockerfile which is original Dockerfile
# who solve the problem ever met in making a user-
# -defined mysql database through "source xx.sql"
# This problem suck me for around 4 days..


FROM maxwell/mysql:v23

MAINTAINER liyao.miao@yeepay.com 

ENV MYSQL_ALLOW_EMPTY_PASSWORD=true

RUN ulimit -n 10240

CMD ["/root/startAll.sh"]

EXPOSE 3306

