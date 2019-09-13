FILE=main

default:
	go build ${FILE}.go
	mv ${FILE} ${FILE}.cgi

clean:
	rm -rf ${FILE}.cgi
