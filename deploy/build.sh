#!/usr/bin/env bash
set -e

echo Building Domio Public...
#rm -rf /domio
#mkdir /domio

cd ~/domiobackoffice

#=====================================================================================================================

buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'`
buildstamp=`date -u '+%s'`
hash=`git rev-parse --short HEAD`
version=`git tag -l --points-at HEAD`

echo ------------------------------------------------------
echo "Buildstamp: ${buildstamp}"
echo "Hash:       ${hash}"
echo "Version:    ${version}"
echo ------------------------------------------------------

go build -o /usr/local/bin/domio_backoffice -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" domio_backoffice

#=====================================================================================================================

#/usr/local/bin/domio_backoffice init --aws-access-key-id=$AWS_ACCESS_KEY_ID \
#                                 --aws-secret-access-key=$AWS_SECRET_ACCESS_KEY \
#                                 --db-name=$DOMIO_DB_NAME \
#                                 --db-user=$DOMIO_DB_USER \
#                                 --db-password=$DOMIO_DB_PASSWORD

/usr/local/bin/domio_backoffice init --env=staging

cd /
rm -rf ~/domiobackoffice

service domio_backoffice stop
service domio_backoffice start

echo Domio Public is built and ready!

logger -n logs5.papertrailapp.com -t deploy -P 18422 -p user.notice "Domio Public is built and ready!"