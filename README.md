up2streamable


# howtouse

## install

on XX pi console
```
go get github.com/m0a/up2streamable
```

OR PC

```
GOOS=linux GOARCH=arm go build github.com/m0a/up2streamable
scp up2streamable pi@192.168.XX.XX:~/bin/

```

## cmd

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

