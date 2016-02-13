#! /bin/bash
# This is the script who can start mysqld 
# and then import the exp.sql
# Attention: the last sleep 60 is so important
# that it won't work without the sentence. 

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
            #break
            exit 1
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

}&
#
#mysqld >/root/a.log 2>&1
#if [ $? -eq 0 ]
#then
#    echo "0 mysql"
#else
#    echo "1 mysql"
#    mysqld >/root/a.log 2&>1
#    if [ $? -eq 0 ]
#    then
#        echo "0 mysql 2"
#    else
#        echo "1 mysql 2"
#    fi
#fi
{
    while true
    do
        sleep 1
        mysqld >/root/a.log 2>&1
        if [ $? -eq 0 ]
        then
            echo "0 mysql"
            #break
            exit 1
        else
            echo "1 mysql"
        fi
    done
}&


while true
do
    sleep 60
done
