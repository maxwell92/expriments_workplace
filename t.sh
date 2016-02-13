#! /bin/bash

touch /root/a.log

mysqld --initialize-insecure --user=root 

{
#    sleep 30
    succ="Version"
    while true
    do
        sleep 1
        grep $succ /root/a.log
        if [ $? -eq 0 ]
        then
            /bin/bash /root/importSql.sh
            echo $? "import"
            break
        fi 
    done
#    echo "30 done"
#    /bin/bash /root/importSql.sh
#    echo "import"
#    echo $?
#    mysql -u root <<EOF
#    create database exp;
#    use exp;
#    source /root/exp.sql;
#    EOF
}

mysqld >/root/a.log 2>&1
echo $? "mysql"
