ARG BASE_IMAGE=alpine:3.10
FROM ${BASE_IMAGE}

ARG SMARTEMS_TGZ="smartems-latest.linux-x64-musl.tar.gz"

COPY ${SMARTEMS_TGZ} /tmp/smartems.tar.gz

# Change to tar xfzv to make tar print every file it extracts
RUN mkdir /tmp/smartems && tar xfz /tmp/smartems.tar.gz --strip-components=1 -C /tmp/smartems

FROM ${BASE_IMAGE}

ARG GF_UID="472"
ARG GF_GID="472"

ENV PATH=/usr/share/smartems/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin \
    GF_PATHS_CONFIG="/etc/smartems/smartems.ini" \
    GF_PATHS_DATA="/var/lib/smartems" \
    GF_PATHS_HOME="/usr/share/smartems" \
    GF_PATHS_LOGS="/var/log/smartems" \
    GF_PATHS_PLUGINS="/var/lib/smartems/plugins" \
    GF_PATHS_PROVISIONING="/etc/smartems/provisioning"

WORKDIR $GF_PATHS_HOME

RUN apk add --no-cache ca-certificates bash tzdata && \
    apk add --no-cache --upgrade --repository=http://dl-cdn.alpinelinux.org/alpine/edge/main openssl musl-utils

# Oracle Support for x86_64 only
RUN if [ `arch` = "x86_64" ]; then \
      apk add --no-cache --upgrade --repository=http://dl-cdn.alpinelinux.org/alpine/edge/community libaio libnsl && \
      ln -s /usr/lib/libnsl.so.2 /usr/lib/libnsl.so.1 && \
      wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.30-r0/glibc-2.30-r0.apk \
        -O /tmp/glibc-2.30-r0.apk && \
      wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.30-r0/glibc-bin-2.30-r0.apk \
        -O /tmp/glibc-bin-2.30-r0.apk && \
      apk add --allow-untrusted /tmp/glibc-2.30-r0.apk /tmp/glibc-bin-2.30-r0.apk && \
      rm -f /tmp/glibc-2.30-r0.apk && \
      rm -f /tmp/glibc-bin-2.30-r0.apk && \
      rm -f /lib/ld-linux-x86-64.so.2 && \
      rm -f /etc/ld.so.cache; \
    fi

# PhantomJS
RUN if [ `arch` = "x86_64" ]; then \
      apk add --no-cache --virtual phantomjs-utils curl && \
      cd /tmp && \
      curl -Ls https://github.com/dustinblackman/phantomized/releases/download/2.1.1/dockerized-phantomjs.tar.gz | tar xz && \
      cp -R lib lib64 / && \
      cp -R usr/lib/x86_64-linux-gnu /usr/lib && \
      cp -R usr/share/fonts /usr/share && \
      cp -R etc/fonts /etc && \
      rm -rf /tmp/* && \
      apk del --no-cache phantomjs-utils; \
    fi

COPY --from=0 /tmp/smartems "$GF_PATHS_HOME"

RUN mkdir -p "$GF_PATHS_HOME/.aws" && \
    addgroup -S -g $GF_GID smartems && \
    adduser -S -u $GF_UID -G smartems smartems && \
    mkdir -p "$GF_PATHS_PROVISIONING/datasources" \
             "$GF_PATHS_PROVISIONING/dashboards" \
             "$GF_PATHS_PROVISIONING/notifiers" \
             "$GF_PATHS_LOGS" \
             "$GF_PATHS_PLUGINS" \
             "$GF_PATHS_DATA" && \
    cp "$GF_PATHS_HOME/conf/sample.ini" "$GF_PATHS_CONFIG" && \
    cp "$GF_PATHS_HOME/conf/ldap.toml" /etc/smartems/ldap.toml && \
    chown -R smartems:smartems "$GF_PATHS_DATA" "$GF_PATHS_HOME/.aws" "$GF_PATHS_LOGS" "$GF_PATHS_PLUGINS" "$GF_PATHS_PROVISIONING" && \
    chmod -R 777 "$GF_PATHS_DATA" "$GF_PATHS_HOME/.aws" "$GF_PATHS_LOGS" "$GF_PATHS_PLUGINS" "$GF_PATHS_PROVISIONING"

EXPOSE 3000

COPY ./run.sh /run.sh

USER smartems
ENTRYPOINT [ "/run.sh" ]
