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
#curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
echo "====== set warehouse ======"
# sudo add-apt-repository \
# "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu \
# $(lsb_release -cs) \
# stable"
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
echo "====== install docker-componse ======"
sudo curl -L "https://github.com/docker/compose/releases/download/1.22.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

