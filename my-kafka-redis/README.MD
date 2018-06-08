git clone https://github.com/edenhill/librdkafka.git
cd librdkafka
./configure --prefix /usr
sudo make
sudo make install

~/.bash_profile
export PKG_CONFIG_PATH=/usr/lib/pkgconfig

yum -y update gcc
yum -y install gcc+ gcc-c++

1. [save dependency to vendor]
$GOPATH/bin/godep save