# Repository updates deb/rpm



## Testing

It's possible to test the repo updates for rpm and deb by running the test scripts within a docker container like this. Tests are being executed by using two buckets on gcp setup for testing.

```bash
docker run -ti --rm -u 0:0 smartems/smartems-ci-deploy:1.2.2 bash # 1.2.2 is the newest image at the time of writing
# in the container:
mkdir -p /dist

#outside of container:
cd <smartems project dir>/..
docker cp smartems <container_name>:/
docker cp <gpg.key used for signing> <container_name>:/private.key

#in container:
./scripts/build/update_repo/load-signing-key.sh
cd dist && wget https://dl.smartems.com/oss/release/smartems_5.4.3_amd64.deb && wget https://dl.smartems.com/oss/release/smartems-5.4.3-1.x86_64.rpm && cd ..

#run these scripts to update local deb and rpm repos and publish them:
./scripts/build/update_repo/test-update-deb-repo.sh <gpg key password>
./scripts/build/update_repo/test-publish-deb-repo.sh
./scripts/build/update_repo/test-update-rpm-repo.sh <gpg key password>
./scripts/build/update_repo/test-publish-rpm-repo.sh

```
