#! /bin/bash
# this script was used for change owner of directory
# /var/run/mysqld (defined in /etc/mysql/my.cnf)
# to mysql:mysql after 15 seconds waiting (wait for mysqld start)

sleep 15
chown -R mysql:mysql /var/run/mysqld
