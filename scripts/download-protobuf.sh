#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

root="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/.."
cd $root

version="3.6.1"
vhash="6003de742ea3fcf703cfec1cd4a3380fd143081a2eb0e559065563496af27807"
protoc="$PWD/download/bin/protoc"
url="https://github.com/protocolbuffers/protobuf/releases/download/v${version}/protoc-${version}-linux-x86_64.zip"

if [ ! -d download ]; then
    mkdir download
fi

if [ ! -d download/bin ]; then
    mkdir download/bin
fi

if [ -f $protoc ]; then
    current=$($protoc --version | awk '{print $2}')
    if [ $current = $version ]; then
        echo "already have protoc v${version}"
        exit 0
    else
        echo "downloading protoc v${version}"
    fi
else
    echo "downloading protoc v${version}"
fi

if [ -d download/tmp ]; then
    rm -rf download/tmp
fi
mkdir download/tmp
cd download/tmp
wget -q -O protoc.zip $url
sha=$(sha256sum  protoc.zip | awk '{print $1}')
if [ $sha != $vhash ]; then
    echo "incorrect hash: $sha"
    echo "expected: $vhash"
    exit 1
fi
unzip protoc.zip > /dev/null
cp bin/protoc ../bin/
cd ..
rm -rf tmp
echo "downloaded protoc v${version}"