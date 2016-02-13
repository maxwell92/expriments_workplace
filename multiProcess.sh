#! /bin/bash
# A script for testing multiprocess in Docker CMD

sleep 23
echo `date` "3 done" > mul.log

{
    sleep 27
    echo `date` "7 done" >> mul.log
}&

{
    sleep 25
    echo `date` "5 done" >> mul.log
}&


sleep 60
