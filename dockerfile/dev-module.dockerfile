
    FROM alpine:3.8 as source
    RUN apk update && apk upgrade && \
        apk add --no-cache bash git openssh
    RUN git clone  https://github.com/Urothis/nwn-order.git 
    ENTRYPOINT ["sleep 10"]

    FROM jakkn/nwn-devbase as modulebuild
    USER root
    COPY --from=source /nwn-order/testing /home/devbase/build

    FROM nwnxee/unified:latest
    COPY --from=modulebuild /home/devbase/build/server/ ./nwn/home/server/