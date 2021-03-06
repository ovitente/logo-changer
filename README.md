![Building App](https://github.com/ovitente/logo-changer/workflows/Building%20App/badge.svg)

<p align="center">
<img src="logo-main.png">
</p>

# logo-changer
Tool for changing banner logo on TeamSpeak3 server.

### How does it work?
*It uses telnet access to reach remote teamspeak server API, called ServerQuery (short SQ).*

* You need to have ftp or http server with directly accessible images
* Create `.env` file using template from this repo or `export` values as env vars. Work in both ways.
* Run logo-changer.
* It randomly chooses logo by filename from local `list-*` files from this repo, and make sure that **both images filenames in list and at your images server must be the same.** Otherwise you will not see desired image by link.
* Control frequency of changing logo in seconds by live editing env var `LOGO_UPDATE_TIMEOUT=5`.

### Where can i deploy it?
* For now it is successfully deployed at heroku. I am controlling it through `heroku` cli app, aslo live changing of env vars
* GCP App Engine
* Any VPS
* Local pc
