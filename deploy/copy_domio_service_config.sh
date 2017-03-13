#!/usr/bin/env bash
set -e

echo Copying Domio Service Config...

yes | cp -rf ~/domiobackoffice/deploy/config/domio_service.sh /etc/init.d/domio_backoffice

systemctl daemon-reload

echo Domio Service Config copied!