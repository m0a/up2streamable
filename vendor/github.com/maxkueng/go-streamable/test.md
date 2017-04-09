Running tool: /Users/m0a/.anyenv/envs/goenv/versions/1.8/bin/go test -timeout 30s -tags  -run ^Test_UploadVideo$

create NewRequest
authenticateHTTPRequest
req.Header.Set Content-Type
start upload.
start write
multipartWriter CreateFormFile.


req => &http.Request{Method:"POST", URL:(*url.URL)(0xc420514080), Proto:"HTTP/1.1", ProtoMajor:1, ProtoMinor:1, Header:http.Header{"Content-Type":[]string{"multipart/form-data; boundary=7cdd6ce13e672ce350d9de8319029e6e786e822c4bf7a3b7090084ce7dc6"}}, Body:ioutil.nopCloser{Reader:(*bytes.Buffer)(0xc4200c0e70)}, GetBody:(func() (io.ReadCloser, error))(0x1234c50), ContentLength:1222787, TransferEncoding:[]string(nil), Close:false, Host:"api.streamable.com", Form:url.Values(nil), PostForm:url.Values(nil), MultipartForm:(*multipart.Form)(nil), Trailer:http.Header(nil), RemoteAddr:"", RequestURI:"", TLS:(*tls.ConnectionState)(nil), Cancel:(<-chan struct {})(nil), Response:(*http.Response)(nil), ctx:context.Context(nil)}



req => &http.Request{Method:"POST", URL:(*url.URL)(0xc42001ec80), Proto:"HTTP/1.1", ProtoMajor:1, ProtoMinor:1, Header:http.Header{"Content-Type":[]string{"multipart/form-data; boundary=5ac1d8e294e1ca308ee67468fa0e3236a9d0671165bc1d0d29e60d9efd33"}}, Body:(*io.PipeReader)(0xc42000e0d0), GetBody:(func() (io.ReadCloser, error))(nil), ContentLength:0, TransferEncoding:[]string(nil), Close:false, Host:"api.streamable.com", Form:url.Values(nil), PostForm:url.Values(nil), MultipartForm:(*multipart.Form)(nil), Trailer:http.Header(nil), RemoteAddr:"", RequestURI:"", TLS:(*tls.ConnectionState)(nil), Cancel:(<-chan struct {})(nil), Response:(*http.Response)(nil), ctx:context.Context(nil)}
io.Copy start
io.Copy end
pipeWriter close
&{400 BAD REQUEST 400 HTTP/1.1 1 1 map[Server:[nginx/1.10.0 (Ubuntu)] Set-Cookie:[session=ZX8G2Z0NW60SOJ4GU5ULNF0Y; Domain=.streamable.com; Expires=Thu, 25-Aug-2044 01:28:36 GMT; Path=/] Content-Length:[18] Access-Control-Allow-Credentials:[true] Retry-After:[1714] Date:[Sun, 09 Apr 2017 01:28:36 GMT] X-Ratelimit-Limit:[60] X-Ratelimit-Remaining:[51] X-Ratelimit-Reset:[1491703031] Connection:[keep-alive] Access-Control-Allow-Origin:[*] Content-Type:[text/html; charset=utf-8]] 0xc4201ee080 18 [] false false map[] 0xc42000b000 0xc42040f130}
--- FAIL: Test_UploadVideo (4.90s)
	assertions.go:226: 
                          
	Error Trace:	streamable_test.go:19
		
	Error:      	Expected nil, but got: &errors.errorString{s:"upload failed StatusCode:400"}
	assertions.go:226: 
                          
	Error Trace:	streamable_test.go:20
		
	Error:      	Should not be zero, but was {0     map[] [] }
FAIL
exit status 1
FAIL	github.com/m0a/up2streamable/vendor/github.com/maxkueng/go-streamable	4.906s
Error: Tests failed.