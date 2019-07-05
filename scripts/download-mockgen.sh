#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

root="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/.."
cd $root

version="91d713334d1e6ea5c8bde6de1b3d11413e686ecd"
vhash="7a81507a2ca4aa7994a0b06b0d7468f3f0a6dc74acd0dffc66fe66e67fbe828d"
mockgen="$PWD/download/bin/mockgen"
url="https://github.com/linuxuser586/mock/archive/${version}.tar.gz"

if [ ! -d download ]; then
    mkdir download
fi

if [ ! -d download/bin ]; then
    mkdir download/bin
fi

if [ -f ${mockgen}.version ] && [ -f ${mockgen} ]; then
    current=$(cat ${mockgen}.version)
    if [ $current = $version ]; then
        echo "already have mockgen v${version}"
        exit 0
    else
        echo "downloading mockgen v${version}"
    fi
else
    echo "downloading mockgen v${version}"
fi

if [ -d download/tmp ]; then
    rm -rf download/tmp
fi
mkdir download/tmp
cd download/tmp
wget -q -O mockgen.tar.gz $url
sha=$(sha256sum  mockgen.tar.gz | awk '{print $1}')
if [ $sha != $vhash ]; then
    echo "incorrect hash: $sha"
    echo "expected: $vhash"
    exit 1
fi
tar xf mockgen.tar.gz
mkdir -p src/github.com/golang
mv mock-${version} src/github.com/golang/mock
GOPATH=${PWD}
go install github.com/golang/mock/mockgen
mv bin/mockgen ../bin
echo $version > ${mockgen}.version
rm -rf tmp
echo "downloaded mockgen v${version}"