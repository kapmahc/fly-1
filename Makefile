dist=build
pkg=github.com/kapmahc/fly/web

VERSION=`git rev-parse --short HEAD`
BUILD_TIME=`date -R`
AUTHOR_NAME=`git config --get user.name`
AUTHOR_EMAIL=`git config --get user.email`
COPYRIGHT=`head -n 1 LICENSE`
USAGE=`sed -n '3p' README.md`

build: backend frontend
	tar jcvf dist.tar.bz2 $(dist)


backend:
	go build -ldflags "-s -w -X ${pkg}.Version=${VERSION} -X '${pkg}.BuildTime=${BUILD_TIME}' -X '${pkg}.AuthorName=${AUTHOR_NAME}' -X ${pkg}.AuthorEmail=${AUTHOR_EMAIL} -X '${pkg}.Copyright=${COPYRIGHT}' -X '${pkg}.Usage=${USAGE}'" -o ${dist}/fly main.go
	-cp -rv locales db templates $(dist)/

frontend:
	cd front && npm run build
	cp -rv front/dist $(dist)/public


clean:
	-rm -rv $(dist) dist.tar.bz2

init:
	govendor sync
	cd front && npm install
