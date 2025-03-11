#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/captcha"
exec "$CURDIR/bin/captcha"
