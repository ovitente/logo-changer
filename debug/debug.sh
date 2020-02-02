#!/bin/bash 
export ARG=$1

export QUOTE="$1"
export ARGUMENTS=( "$@" )
echo $MSG[$@]
echo -e "\nRemoving first argument"
# array=( "${array[@]/$delete}" )
delete=($1)
echo "${ARGUMENTS[@]/$delete}"
ARGUMENTS=( "${ARGUMENTS[@]/$delete}" )

echo -e "\nArray result"
echo ${ARGUMENTS[@]}

echo -e "\nMSG TO SERVER:"
export MSG=${ARGUMENTS[@]}
echo "$MSG"
MSG="${MSG// /\\s}"
echo $MSG
echo " "

CREDENTIALS_FILE="creds.cfg"

function GetCredentials {
    if [ -e $CREDENTIALS_FILE ]; then
        source $CREDENTIALS_FILE
    else
        echo "Cannot find credentials file."
    fi
}

function ChangeLogo {
    expect -c '
        spawn telnet $env(SQUERY_URL) $env(SQUERY_PORT)
        send "login $env(SQUERY_USER) $env(SQUERY_PASSWORD)\r\n"
        send "use sid=$env(SQUERY_SID)\r\n"
        send "serveredit sid=$env(SQUERY_SID) virtualserver_hostbanner_gfx_url=http://$env(FTP_HOST)/quotes/q$env(QUOTE).png\r\n"
		send "quit\r\n"
		interact
'
}

function SendServerMsg {
  MsgWrapper $MSG
  echo "- Sending message [ $1 ]"
    expect -c '
        spawn telnet $env(SQUERY_URL) $env(SQUERY_PORT)
        send "login $env(SQUERY_USER) $env(SQUERY_PASSWORD)\r\n"
        send "use sid=$env(SQUERY_SID)\r\n"
        send "sendtextmessage targetmode=3 target=$env(SQUERY_SID) msg=$env(MSG)\r\n"

		send "quit\r\n"
		interact
'
}

function MsgWrapper {
  MSG="[color=green][b]$1[/b][/color]"
}

GetCredentials
case $1 in
  "sm")
    SendServerMsg $MSG
    ;;
  "change-logo")
    ChangeLogo $1
    ;;
  *)
    echo -e "Commands available: \n
    sm - Send text message to server
    change-logo
    \n"
esac
  


