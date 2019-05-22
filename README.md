# gozer

Fork for personal use.

## Installation & Usage

1. Install from source (`go get github.com/step21/gozer`) **or** download a [release](https://github.com/step21/gozer/releases) for your platform and put it in your PATH.
2. Put your token in `~/.gozer-token` (linux/osx) or in `%USERPROFILE5\.gozer-token` (windows).
3. Run `gozer` to display all networks and clients, use `gozer --online` to only show online clients.

```bash
$ gozer
1234567890abcdef network1 
     	mba_osx                   1111111111 [10.242.52.135]	
     	mba_windows               2222222222 [10.242.190.139]	 Offline
     	pizerow                   3333333333 [10.242.34.69]	
     	raspberrypi               4444444444 [10.242.227.38]	

fedcba0987654321 network2 
     	build_server              7777777777 [192.168.193.2]	
     	mba_osx                   1111111111 [192.168.193.148]	
     	mba_windows               2222222222 [192.168.193.26]	 Offline
     	science                   5555555555 [192.168.193.218]	
     	thewolfgang               6666666666 [192.168.193.66]
```
