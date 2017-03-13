#!/usr/bin/env bash

set -e

logger -n logs5.papertrailapp.com -t deploy -P 18422 -p user.notice "Domio Backoffice deploy has started..."

cd ~/domiobackoffice
export GOPATH=$PWD
echo $GOPATH

sh ~/domiobackoffice/deploy/apt_update.sh

if ! [ -x "$(command -v go)" ]; then
   echo 'go is not installed.' >&2
   sh ~/domiobackoffice/deploy/install_go.sh
  else
   echo "Go is already installed!" >&2
fi

if ! [ -x "$(command -v letsencrypt)" ]; then
   echo 'letsencrypt is not installed.' >&2
   sh ~/domiobackoffice/deploy/install_letsencrypt.sh
  else
   echo "letsencrypt is already installed!" >&2
fi

if ! [ -x "$(command -v nginx)" ]; then
   echo 'nginx is not installed.' >&2
   sh ~/domiobackoffice/deploy/install_nginx.sh
  else
   echo "nginx is already installed!" >&2
fi

sh ~/domiobackoffice/deploy/install_deps.sh
sh ~/domiobackoffice/deploy/copy_domio_service_config.sh
sh ~/domiobackoffice/deploy/copy_nginx_config_files.sh

sh ~/domiobackoffice/deploy/build.sh

service domio_backoffice restart

cd /