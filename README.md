![Building App](https://github.com/ovitente/logo-changer/workflows/Building%20App/badge.svg)
![logo](logo-main.png)
# logo-changer
Tool for changing banner logo on TeamSpeak3 server by sending command through Telnet.

---
Debug command:  
```
serveredit sid=$env(SQUERY_SID) virtualserver_hostbanner_gfx_url=http://$env(FTP_URL)/quotes/$env(RDY_IMG)\r\n"
```
