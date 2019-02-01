
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
         NWNX_REDIS_PUBSUB_CHANNELS=Heartbeat,Discord:Out \
         # Environmental Variables
         # Order
         NWN_ORDER_PORT=5750 \
         NWN_ORDER_MODULE_NAME=Order \ #GetModuleName(); || o_inc.nss:Line 1 
         # Redis
         NWN_ORDER_REDIS_PORT=6379 \
         # API
         NWN_ORDER_WEBHOOKS_GITHUB= \
         # Heartbeat
         NWN_ORDER_PLUGIN_HEARTBEAT_ENABLED=1 \
         NWN_ORDER_PLUGIN_HEARTBEAT_VERBOSE=0 \
         NWN_ORDER_PLUGIN_HEARTBEAT_ONE_MINUTE=1 \
         NWN_ORDER_PLUGIN_HEARTBEAT_FIVE_MINUTE=1 \
         NWN_ORDER_PLUGIN_HEARTBEAT_THIRTY_MINUTE=1 \
         NWN_ORDER_PLUGIN_HEARTBEAT_ONE_HOUR=1 \
         NWN_ORDER_PLUGIN_HEARTBEAT_SIX_HOUR=1 \
         NWN_ORDER_PLUGIN_HEARTBEAT_TWELVE_HOUR=1 \
         NWN_ORDER_PLUGIN_HEARTBEAT_TWENTYFOUR_HOUR=1 \
         # Logging
         NWN_ORDER_PLUGIN_LOG_ENABLED=1 \
         NWN_ORDER_PLUGIN_LOG_TARGET=wip