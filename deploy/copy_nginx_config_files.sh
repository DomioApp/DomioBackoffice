#!/usr/bin/env bash
set -e

echo Copying nginx config files...

yes | cp -rf ~/domiopublic/deploy/config/config.nginx /etc/nginx/nginx.conf

if [ -f "/run/nginx.pid" ]
then
    service nginx reload
else
	service nginx start
fi


echo nginx config files copied!