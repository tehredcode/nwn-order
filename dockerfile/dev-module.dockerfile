FROM alpine:3.8 as source
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN git clone  https://github.com/Urothis/nwn-order.git && \
    cd nwn-order && \
    git checkout dev 
ENTRYPOINT ["sleep 10"]
FROM jakkn/nwn-devbase as modulebuild
USER root
COPY --from=source /nwn-order/testing /home/devbase/build
FROM nwnxee/unified:latest
COPY --from=modulebuild /home/devbase/build/server/ ./nwn/home/server/
ENV  NWN_MODULE=Order \
     NWN_SERVERNAME=Order \
     NWN_PUBLICSERVER=1 \
     NWN_MAXCLIENTS=255 \
     NWN_MINLEVEL=1 \
     NWN_MAXLEVEL=40 \
     NWN_PAUSEANDPLAY=0 \
     NWN_PVP=2 \
     NWN_SERVERVAULT=1 \
     NWN_ELC=0 \
     NWN_ILR=0 \
     NWN_GAMETYPE=0 \
     NWN_ONEPARTY=0 \
     NWN_DIFFICULTY=3 \
     NWN_AUTOSAVEINTERVAL=0 \
     NWN_RELOADWHENEMPTY=0 \
     NWN_PLAYERPASSWORD= \
     NWN_DMPASSWORD=123 \
     NWN_ADMINPASSWORD= \
     NWNX_CORE_LOG_LEVEL=6 \
     NWNX_ADMINISTRATION_SKIP=n \
     NWNX_REDIS_SKIP=n \
     NWNX_TIME_SKIP=n \
     NWNX_WEBHOOK_SKIP=n \
     NWNX_REDIS_HOST=redis \
     NWNX_REDIS_PORT=6379 \
     NWNX_REDIS_PUBSUB_CHANNELS=Heartbeat,Discord:In \
     NWN_ORDER_PORT=5750 \
     NWN_ORDER_MODULE_NAME=order \
     NWN_ORDER_REDIS_HOST=redis \
     NWN_ORDER_REDIS_PORT=6379 \
     NWN_ORDER_REDIS_PUBSUB_VERBOSE=1 \
     NWN_ORDER_WEBHOOKS=0 \
     NWN_ORDER_GITHUB_WEBHOOK_SECRET= \
     NWN_ORDER_GITLAB_WEBHOOK_SECRET= \
     NWN_ORDER_DISCORD_ENABLED=0 \
     NWN_ORDER_DISCORD_BOT_KEY= \
     NWN_ORDER_DISCORD_PUBLIC_BOT_ROOM= \
     NWN_ORDER_DISCORD_PRIVATE_BOT_ROOM= \
     NWN_ORDER_HEARTBEAT_ENABLED=1 \
     NWN_ORDER_HEARTBEAT_ONE_MINUTE=1 \
     NWN_ORDER_HEARTBEAT_FIVE_MINUTE=1 \
     NWN_ORDER_HEARTBEAT_THIRTY_MINUTE=1 \
     NWN_ORDER_HEARTBEAT_ONE_HOUR=1 \
     NWN_ORDER_HEARTBEAT_SIX_HOUR=1 \
     NWN_ORDER_HEARTBEAT_TWELVE_HOUR=1 \
     NWN_ORDER_HEARTBEAT_TWENTYFOUR_HOUR=1 \
     NWN_ORDER_LOG_ENABLED=1 \
     NWN_ORDER_LOG_TARGET=wip