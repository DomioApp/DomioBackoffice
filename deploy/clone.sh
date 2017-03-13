#!/usr/bin/env bash

cd ~/
rm -rf ~/domiopublic
git clone git@gitlab.com:basharov/DomioPublic.git ~/domiopublic
cd ~/domiopublic
git tag -l --points-at HEAD
