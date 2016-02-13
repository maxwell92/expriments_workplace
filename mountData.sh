#! /bin/bash
# tried to import the database exp.sql by overwritten
# the filepath /var/lib/mysql. It doesn't work now and
# we can see the database imported successfully but can't
# find the right table. It needs to be considered later.

cp -r -f mysql /var/lib/
chown -R mysql:mysql /var/lib/mysql/exp
