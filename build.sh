#!/bin/bash

set -e

# working directory must clean
if [ $(git status | grep -c clean) -eq 0 ]; then
    git status
    echo "commit all changes"
    exit 1
fi

# git-buildpackages

maintainer='Michael Rennecke <michael.rennecke@gmail.com>'
vendor="$maintainer"
url="https://github.com/0rph3us/sensorserver"
name=sensorserver
version="0.0.1"

# build ARMv7 package
env GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-X main.Version=$version -X main.BuildTime=$(date -u '+%Y-%m-%d_%H:%M:%S_UTC') -X main.Commit=$(git rev-parse HEAD)" cmd/main.go

# build deb package
[ -f *.deb ] && rm *.deb

mkdir -p package/usr/sbin/
mkdir -p package/etc/sensorserver/
mkdir -p package/var/lib/sensorserver/

mv main package/usr/sbin/sensorserver
cp config.toml package/etc/sensorserver/config.toml
cp -r templates package/etc/sensorserver/
cp -r assets package/var/lib/sensorserver/
cp -r resources package/var/lib/sensorserver/



fpm -s dir -t deb -n $name \
  -C package \
  --vendor "$vendor" \
  --maintainer "$maintainer" \
  --version "$version" \
  --category misc \
  --url "$url" \
  --license "GNU GENERAL PUBLIC LICENSE Version 2, June 1991" \
  --architecture armv7l \
  --after-install package/postinstall \
  -d adduser \
  etc lib usr var
