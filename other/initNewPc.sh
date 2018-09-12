echo "====== Glad to start the initialization environment for you ======"
echo "====== Please confirm it is over the wall ======

echo "====== start Admin rules and entry your is password ======"
sudo su
echo "====== install gvm dependent ======"
echo "====== entry continue ======"
sudo apt-get install curl git mercurial make binutils bison gcc build-essential

echo "====== install gvm ======"
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

echo "====== source gvm path ====="
source /root/.gvm/scripts/gvm

echo "====== install golang ======="
gvm install go1.10.2 -B

echo "====== set golang version ======"
gvm use go1.10.2 --default

echo "========================================================"
echo "==================== install docker ===================="
echo "========================================================"

echo "====== uninstall docker in old ======"
sudo apt-get remove docker docker-engine docker.io
echo "====== update apt package ======"
sudo apt-get update
echo "====== set apt https ======"
sudo apt-get install \
apt-transport-https \
ca-certificates \
curl \
software-properties-common
echo "====== add key ======"
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -ok
echo "====== set warehouse ======"
sudo add-apt-repository \
"deb [arch=amd64] https://download.docker.com/linux/ubuntu \
$(lsb_release -cs) \
stable"


echo "====== apt update ===="
sudo apt-get update

echo "====== install new version docker-CE ======"
sudo apt-get install docker-ce

echo "====== docker run ======"
sudo docker run hello-world
















