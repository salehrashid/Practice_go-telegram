FROM linuxserver/wireguard

# Install envsubst to replace template
RUN apk add --no-cache gettext

# Copy custom entrypoint
COPY wg-config/wg-confs/wg0.conf /etc/wireguard/wg0.conf
COPY docker/wireguard/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
