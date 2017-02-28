up2streamable

[![wercker status](https://app.wercker.com/status/d0ad96a0302bfc42785b135aace16d00/m/ "wercker status")](https://app.wercker.com/project/byKey/d0ad96a0302bfc42785b135aace16d00)


# howtouse

## install

on raspberrypi or orangepi console
```
$ wget  https://github.com/m0a/up2streamable/releases/download/0.1/up2streamable
```

OR 

on PC

```
$ GOOS=linux GOARCH=arm go build github.com/m0a/up2streamable

$ scp up2streamable pi@192.168.XX.XX:~/bin/

```

## command line example

```
up2streamable -path ~/.octoprint/timelapse -delete -u username -p password
```

## options
```
up2streamable --help
Usage of up2streamable:
  -delete
    	delete when upload success
  -e string
    	file extension specify (default ".mpg")
  -p string
    	password
  -path string
    	root file path (default "./")
  -u string
    	username
```

