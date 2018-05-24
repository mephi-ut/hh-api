Installation
============

Just a proof of work (for Debian):

```sh
apt-get update
apt-get install -y curl apt-transport-https gnupg libc6-dev
apt-get install curl python-software-properties
curl -sL https://deb.nodesource.com/setup_10.x | base -
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
apt-get update
apt-get install -y golang yarn sudo nodejs

useradd -m site
su -l site <<EOF
go get github.com/mephi-ut/hh-api
#go get github.com/xaionaro/reform/reform
#go install github.com/xaionaro/reform/reform
cd /home/site/go/src/github.com/mephi-ut/hh-api
~/go/bin/reform models
#go get github.com/mephi-ut/hh-api
go install github.com/mephi-ut/hh-api
cd frontend
yarn
yarn build
EOF

apt-get install -y nginx
rm -f /etc/nginx/sites-enabled/default
ln -s /home/site/go/src/github.com/mephi-ut/hh-api/doc/nginx.conf /etc/nginx/sites-enabled/hh-api

# let's encrypt stuff here: apt-get install -y certbot && certbot certonly

nginx -s reload

sudo -u site /home/go/bin/hh-api
```

