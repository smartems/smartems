#!/bin/sh

# no relation to publish.go

# shellcheck disable=SC2124

EXTRA_OPTS="$@"

# Right now we hack this in into the publish script.
# Eventually we might want to keep a list of all previous releases somewhere.
_releaseNoteUrl="https://community.smartems.com/t/release-notes-v6-5-x/22704"
_whatsNewUrl="https://smartems.com/docs/guides/whats-new-in-v6-5"

./scripts/build/release_publisher/release_publisher \
    --wn "${_whatsNewUrl}" \
    --rn "${_releaseNoteUrl}" \
    --version "${CIRCLE_TAG}" \
    --apikey  "${SMARTEMS_COM_API_KEY}" "${EXTRA_OPTS}"
