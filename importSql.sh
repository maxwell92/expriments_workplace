#! /bin/bash

#service mysql restart

mysql -u root <<EOF
create database exp;
use exp;
source /root/exp.sql;
EOF

mysql -u root <<EOF
use mysql;
create user maxwell identified by "";
grant all on *.* to maxwell;
flush privileges;
EOF

