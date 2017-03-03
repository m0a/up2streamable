#up2streamable

[![wercker status](https://app.wercker.com/status/d0ad96a0302bfc42785b135aace16d00/m/ "wercker status")](https://app.wercker.com/project/byKey/d0ad96a0302bfc42785b135aace16d00)



## install

on raspberrypi or orangepi console
```
cd ~/bin/
wget  https://github.com/m0a/up2streamable/releases/download/0.1/up2streamable
chmod +x up2streamable
```

OR

on PC
```
GOOS=linux GOARCH=arm go build github.com/m0a/up2streamable
scp up2streamable pi@192.168.XX.XX:~/bin/
```

## command line example

```
up2streamable -path ~/.octoprint/timelapse -delete -u username -p password
```

## options
```
Usage of up2streamable:
  -delete
    	delete when upload success
  -e string
    	specify file extension (default ".mpg")
  -p string
    	password
  -path string
    	search file path (default "./")
  -timeout int
    	specify file upload timeout(minutes) (default 10)
  -u string
    	username
```

