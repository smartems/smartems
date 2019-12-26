ARG BASE_IMAGE=ubuntu:18.10
FROM ${BASE_IMAGE} AS smartems-builder

ARG SMARTEMS_TGZ="smartems-latest.linux-x64.tar.gz"

COPY ${SMARTEMS_TGZ} /tmp/smartems.tar.gz

RUN mkdir /tmp/smartems && tar xfz /tmp/smartems.tar.gz --strip-components=1 -C /tmp/smartems

FROM ${BASE_IMAGE}

EXPOSE 3000

# Set DEBIAN_FRONTEND=noninteractive in environment at build-time
ARG DEBIAN_FRONTEND=noninteractive
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

# Install dependencies
# We need curl in the image, and if the architecture is x86-64, we need to install libfontconfig1 for PhantomJS
RUN if [ `arch` = "x86_64" ]; then \
        apt-get update && apt-get upgrade -y && apt-get install -y ca-certificates libfontconfig1 curl && \
        apt-get autoremove -y && rm -rf /var/lib/apt/lists/*; \
    else \
        apt-get update && apt-get upgrade -y && apt-get install -y ca-certificates curl && \
        apt-get autoremove -y && rm -rf /var/lib/apt/lists/*; \
    fi

COPY --from=smartems-builder /tmp/smartems "$GF_PATHS_HOME"

RUN mkdir -p "$GF_PATHS_HOME/.aws" && \
    addgroup --system --gid $GF_GID smartems && \
    adduser --system --uid $GF_UID --ingroup smartems smartems && \
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

COPY ./run.sh /run.sh

USER smartems
ENTRYPOINT [ "/run.sh" ]
