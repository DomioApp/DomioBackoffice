#!/usr/bin/env bash

cd ~/
rm -rf ~/domiobackoffice
git clone git@gitlab.com:basharov/domiobackoffice.git ~/domiobackoffice
cd ~/domiobackoffice
git tag -l --points-at HEAD
