#!/bin/bash
set -e
set -o pipefail

cgogen ykpiv.yml
cp /usr/src/yubico-piv-tool/lib/*.h .
mv ykpiv/* .
rm -rf ykpiv
