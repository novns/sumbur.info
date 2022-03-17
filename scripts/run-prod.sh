#! /bin/sh

set -e


cd "$(dirname "$0")/../sumbur"

CGO_ENABLED=0  go build  -o ../public/sumbur  -v

strip -s ../public/sumbur


cd ../public/static

gzip -fkv --  *.css  *.js


cd ../..

rsync \
    --archive \
    --delete-excluded \
    --exclude "sumbur-debug*" \
    --verbose \
    public/ \
    sumbur.info:/srv/web/sumbur.info/sumbur/
