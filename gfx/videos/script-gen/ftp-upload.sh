#!/bin/bash
#$1 is the file name
#usage:this_script <filename>
HOST="q.rs.net.ua"
USER="$1"
PASSWD="$2"
FILE="novogodnee-pozdravlenie-im-danilova-2021.png"
REMOTEPATH='/videos'

ftp -n $HOST <<END_SCRIPT
quote USER $USER
quote PASS $PASSWD
cd $REMOTEPATH
put $FILE
quit
END_SCRIPT
exit 0
