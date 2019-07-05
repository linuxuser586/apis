#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

root="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/.."
cd $root

protoc="$PWD/download/bin/protoc"
mockgen="$PWD/download/bin/mockgen"

for proto in $(find protos/ -name "*.proto")
do
    echo "generating proto $proto"
    $protoc --gogofaster_out=plugins=grpc:$GOPATH/src $proto
done

for pb in $(find grpc/ -name "*.pb.go")
do
    echo "generating mock for $pb"
    dest=$(dirname $pb)/mock/$(basename ${pb/.pb./.pb_mock.})
    $mockgen -source=$pb -destination=$dest -package=mock
done
