CRTDIR=$(cd `dirname $0`; pwd)

goxc -os="linux darwin windows freebsd openbsd" -arch="amd64 arm" -n=ImageSplitter -pv=v1.0.1 -wd=${CRTDIR}/../src -d=${CRTDIR}/./release -include=*.go,README*,LICENSE*