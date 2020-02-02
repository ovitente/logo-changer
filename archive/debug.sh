#!/bin/bash 
clear

export QUOTE=$1

CREDENTIALS_FILE="creds.cfg"

function GetCredentials {
    if [ -e $CREDENTIALS_FILE ]; then
        source $CREDENTIALS_FILE
    else
        echo "Cannot find credentials file."
    fi
}

function Connect {
    expect -c '
        spawn telnet $env(SQUERY_URL) $env(SQUERY_PORT)
        send "login $env(SQUERY_USER) $env(SQUERY_PASSWORD)\r\n"
        send "use sid=$env(SQUERY_SID)\r\n"
        send "serveredit sid=$env(SQUERY_SID) virtualserver_hostbanner_gfx_url=http://$env(FTP_HOST)/quotes/q$env(QUOTE).png\r\n"
		send "quit\r\n"
		interact
'
}

GetCredentials
Connect

echo "END OF TELNET SESSION :)"
