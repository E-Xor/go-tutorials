# go-tutorials

## Docker

### Basics

```
alias docker-start='docker-machine start default'
alias docker-eval='eval "$(docker-machine env default)"'
alias go-tutorials-run='docker run -it -v ~/other-apps/go-tutorials/:/root/mounts/go-tutorials -p 80:80 --add-host=mysql_host:$(ipconfig getifaddr en0) --dns=8.8.8.8 go-tutorials:latest bash'

docker-start
docker-eval
docker pull ubuntu
docker run -it ubuntu /bin/bash

root@5b263ef19600:/# apt-get update
root@5b263ef19600:/# apt-get install apache2 -y
root@5b263ef19600:/# exit
docker commit -m "Start go-tutorials" 5b263ef19600 go-tutorials:latest
docker tag c19ac7a16312 winmo.vm:latest
docker images # just to check

go-tutorials-run

apt-get install wget nano -y
wget https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.7.1.linux-amd64.tar.gz
nano /etc/profile
  export PATH=$PATH:/usr/local/go/bin
  export GOPATH=/root/mounts/go-tutorials
source /etc/profile

(exit, docker commit -m "Install Go" 791aaa877bc8 go-tutorials:latest)

go install github.com/E-Xor/hello # path after go-tutorials/src
$GOPATH/bin/hello

# For lib, not executable
go build github.com/E-Xor/stringutil
```
