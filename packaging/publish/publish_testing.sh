#! /usr/bin/env bash
deb_ver=5.1.0-beta1
rpm_ver=5.1.0-beta1

wget https://s3-us-west-2.amazonaws.com/smartems-releases/release/smartems_${deb_ver}_amd64.deb

package_cloud push smartems/testing/debian/jessie smartems_${deb_ver}_amd64.deb
package_cloud push smartems/testing/debian/wheezy smartems_${deb_ver}_amd64.deb
package_cloud push smartems/testing/debian/stretch smartems_${deb_ver}_amd64.deb

wget https://s3-us-west-2.amazonaws.com/smartems-releases/release/smartems-${rpm_ver}.x86_64.rpm

package_cloud push smartems/testing/el/6 smartems-${rpm_ver}.x86_64.rpm
package_cloud push smartems/testing/el/7 smartems-${rpm_ver}.x86_64.rpm

rm smartems*.{deb,rpm}
