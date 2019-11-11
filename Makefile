FILE=main
FOLD=public

default:
	go build ${FILE}.go
	mv ${FILE} ${FOLD}/${FILE}.cgi

run: default
	docker-compose -f docker/docker-compose.yaml build
	docker-compose -f docker/docker-compose.yaml up -d
	echo go build -ldflags '-s -w' -o public/main.cgi main.go
	echo exit
	docker-compose -f docker/docker-compose.yaml exec go bash
	open http://localhost:8080

stop:
	docker-compose -f docker/docker-compose.yaml down

clean:
	rm -rf ${FOLD}/${FILE}.cgi
