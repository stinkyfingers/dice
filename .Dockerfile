$ docker run ubuntu:14.04 grep -v '^#' /etc/apt/sources.list

deb http://archive.ubuntu.com/ubuntu/ trusty main restricted
deb-src http://archive.ubuntu.com/ubuntu/ trusty main restricted

deb http://archive.ubuntu.com/ubuntu/ trusty-updates main restricted
deb-src http://archive.ubuntu.com/ubuntu/ trusty-updates main restricted

deb http://archive.ubuntu.com/ubuntu/ trusty universe
deb-src http://archive.ubuntu.com/ubuntu/ trusty universe
deb http://archive.ubuntu.com/ubuntu/ trusty-updates universe
deb-src http://archive.ubuntu.com/ubuntu/ trusty-updates universe


deb http://archive.ubuntu.com/ubuntu/ trusty-security main restricted
deb-src http://archive.ubuntu.com/ubuntu/ trusty-security main restricted
deb http://archive.ubuntu.com/ubuntu/ trusty-security universe
deb-src http://archive.ubuntu.com/ubuntu/ trusty-security universe