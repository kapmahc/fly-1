# fly

A complete open source e-commerce solution for the Go language(STILL IN DEVELOPMENT).

## Usage

```bash
go get -u github.com/kapmahc/fly
cd $GOPATH/src/github.com/kapmahc/fly
# download packages
make init
# build
make
ls dist
```

## Devleopment

```bash
cd $GOPATH/src/github.com/kapmahc/fly
# generate config.toml
go run main.go g c
# generate nginx.conf
go run main.go g ng
sudo ln -s $GOPATH/src/github.com/kapmahc/fly/etc/nginx/sites-enabled/localhost.conf /etc/nginx/sites-enabled/
sudo nginx -s reload
# start server
./run.sh
```

## Create database

```bash
psql -U postgres
CREATE DATABASE db-name WITH ENCODING = 'UTF8';
CREATE USER user-name WITH PASSWORD 'change-me';
GRANT ALL PRIVILEGES ON DATABASE db-name TO user-name;
```

## Issues

- Rabbitmq Management Plugin(<http://localhost:15612>)

  ```bash
  rabbitmq-plugins enable rabbitmq_management
  rabbitmqctl add_user test test
  rabbitmqctl set_user_tags test administrator
  rabbitmqctl set_permissions -p / test ".*" ".*" ".*"
  ```

- "RPC failed; HTTP 301 curl 22 The requested URL returned error: 301"

  ```bash
  git config --global http.https://gopkg.in.followRedirects true
  ```

- 'Peer authentication failed for user', open file "/etc/postgresql/9.5/main/pg_hba.conf" change line:

  ```
  local   all             all                                     peer  
  TO:
  local   all             all                                     md5
  ```

- Generate openssl certs

  ```bash
  openssl genrsa -out www.change-me.com.key 2048
  openssl req -new -x509 -key www.change-me.com.key -out www.change-me.com.crt -days 3650 # Common Name:*.change-me.com
  ```

- Generate sitemap.xml.gz everyday

  ```bash
  @daily cd /var/www/www.change-me.com && ./fly seo
  ```
- [For gmail smtp](http://stackoverflow.com/questions/20337040/gmail-smtp-debug-error-please-log-in-via-your-web-browser)
