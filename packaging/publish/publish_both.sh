#! /usr/bin/env bash
version=5.4.2

# wget https://dl.smartems.com/oss/release/smartems_${version}_amd64.deb
#
# package_cloud push smartems/stable/debian/jessie smartems_${version}_amd64.deb
# package_cloud push smartems/stable/debian/wheezy smartems_${version}_amd64.deb
# package_cloud push smartems/stable/debian/stretch smartems_${version}_amd64.deb
#
# package_cloud push smartems/testing/debian/jessie smartems_${version}_amd64.deb
# package_cloud push smartems/testing/debian/wheezy smartems_${version}_amd64.deb --verbose
# package_cloud push smartems/testing/debian/stretch smartems_${version}_amd64.deb --verbose

wget https://dl.smartems.com/oss/release/smartems-${version}-1.x86_64.rpm

package_cloud push smartems/testing/el/6 smartems-${version}-1.x86_64.rpm --verbose
package_cloud push smartems/testing/el/7 smartems-${version}-1.x86_64.rpm --verbose

package_cloud push smartems/stable/el/7 smartems-${version}-1.x86_64.rpm --verbose
package_cloud push smartems/stable/el/6 smartems-${version}-1.x86_64.rpm --verbose

rm smartems*.{deb,rpm}
