#!/usr/bin/env bash

set -e

logger -n logs5.papertrailapp.com -t deploy -P 18422 -p user.notice "Domio Public deploy has started..."

cd ~/domiopublic
export GOPATH=$PWD
echo $GOPATH

sh ~/domiopublic/deploy/apt_update.sh

if ! [ -x "$(command -v go)" ]; then
   echo 'go is not installed.' >&2
   sh ~/domiopublic/deploy/install_go.sh
  else
   echo "Go is already installed!" >&2
fi

if ! [ -x "$(command -v letsencrypt)" ]; then
   echo 'letsencrypt is not installed.' >&2
   sh ~/domiopublic/deploy/install_letsencrypt.sh
  else
   echo "letsencrypt is already installed!" >&2
fi

if ! [ -x "$(command -v nginx)" ]; then
   echo 'nginx is not installed.' >&2
   sh ~/domiopublic/deploy/install_nginx.sh
  else
   echo "nginx is already installed!" >&2
fi

sh ~/domiopublic/deploy/install_deps.sh
sh ~/domiopublic/deploy/copy_domio_service_config.sh
sh ~/domiopublic/deploy/copy_nginx_config_files.sh

sh ~/domiopublic/deploy/build.sh

service domio_public restart

cd /