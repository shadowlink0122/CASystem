FILE=demo

default:
	go build ${FILE}.go
	mv ${FILE} ${FILE}.cgi


