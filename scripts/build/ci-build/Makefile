VERSION="dev"
TAG="smartems/build-container"
USER_ID=$(shell id -u)
GROUP_ID=$(shell id -g)

all: build deploy

build:
	docker build -t "${TAG}:${VERSION}" .

deploy:
	docker push "${TAG}:${VERSION}"

run:
	docker run -ti \
		-e "CIRCLE_BRANCH=local" \
		-e "CIRCLE_BUILD_NUM=472" \
		${TAG}:${VERSION} \
		bash

run-with-local-source-live:
	docker run -d \
		-e "CIRCLE_BRANCH=local" \
		-e "CIRCLE_BUILD_NUM=472" \
		-w "/go/src/github.com/smartems/smartems" \
		--name smartems-build \
		-v "${GOPATH}/src/github.com/smartems/smartems:/go/src/github.com/smartems/smartems" \
		${TAG}:${VERSION} \
		bash -c "/tmp/bootstrap.sh; mkdir /.cache; chown "${USER_ID}:${GROUP_ID}" /.cache; tail -f /dev/null"
	docker exec -ti --user "${USER_ID}:${GROUP_ID}" smartems-build bash

run-with-local-source-copy:
	docker run -d \
		-e "CIRCLE_BRANCH=local" \
		-e "CIRCLE_BUILD_NUM=472" \
		-w "/go/src/github.com/smartems/smartems" \
		--name smartems-build \
		${TAG}:${VERSION} \
		bash -c "/tmp/bootstrap.sh; tail -f /dev/null"
	docker cp "${GOPATH}/src/github.com/smartems/smartems" smartems-build:/go/src/github.com/smartems/
	docker exec -ti smartems-build bash

update-source:
	docker cp "${GOPATH}/src/github.com/smartems/smartems" smartems-build:/go/src/github.com/smartems/	

attach:
	docker exec -ti smartems-build bash

attach-live:
	docker exec -ti --user "${USER_ID}:${GROUP_ID}" smartems-build bash

stop:
	docker kill smartems-build
	docker rm smartems-build
