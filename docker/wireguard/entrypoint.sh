#!/bin/sh
set -e

# render wg0.conf dari template
envsubst < /wg-config/wg-confs/wg0.conf > /wg-config/wg0.conf

# pastikan folder /etc/wireguard ada
mkdir -p /etc/wireguard

# copy ke lokasi default yang dicari wg-quick
cp /wg-config/wg0.conf /etc/wireguard/wg0.conf

# jalankan wg-quick
wg-quick up wg0

tail -f /dev/null

exec "$@"
