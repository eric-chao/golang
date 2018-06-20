# librdkafka安装

git clone https://github.com/edenhill/librdkafka.git

cd librdkafka

./configure --prefix /usr
sudo make
sudo make install

~/.bash_profile
export PKG_CONFIG_PATH=/usr/lib/pkgconfig

# librdkafka依赖
yum -y update gcc
yum -y install gcc+ gcc-c++

ldconfig
vim /etc/ld.so.conf
/usr/local/lib
ldconfig

1. [save dependency to vendor]
$GOPATH/bin/godep save

2.
all             ignore-days         expire-days         useLogLog=false
api             ignore-days         ignore-days + 1
daily
hourly
monthly(uv)
stat

[test]
1. go test xxx.go -test.v
2. go test -bench Benchmark_pfadd* -count 10 -cpu 4
   go test -bench Benchmark_hinc* -count 2 -cpu 1,2,3,4

