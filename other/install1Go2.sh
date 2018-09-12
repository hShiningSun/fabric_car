echo "====== Glad to start the initialization environment for you ======"
echo "====== Please confirm it is over the wall ======"

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

echo "====== if ERROR: Failed to download binary go ======"
echo "====== git clone https://github.com/hShiningSun/fabric_installed_file.git ======"

echo "====== export PATH=$PATH:$GO_INSTALL_DIR/go/bin ======"
echo "====== zhixing  go ======"
echo "====== cd root vim .project set GOPATH AND PATH ======"
