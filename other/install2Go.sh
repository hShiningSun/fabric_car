
echo "===== install go1.10.2 ====="
#cp -Rf other/goinstallfile/go1.10.2.linux-amd64.tar.gz /usr/local
# cp -Rf other/goinstallfile/go /usr/local
(
    cd /usr/local
    curl -O https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.10.2.linux-amd64.tar.gz
)

echo " 创建 go 工作目录 src\pak\bin"
echo " .profile 追加PATH=/usr/local/go/bin"
echo " .bashrc  末尾 export GOPATH=/home/car/carGo"
echo "             export GOBIN=$GOPATH/bin"
echo "/root/.bashrc  末尾 PATH=/usr/local/go/bin"
echo "/root/.bashrc  末尾 export GOPATH=/home/car/carGo"
echo "/root/.bashrc  末尾 export GOBIN=$GOPATH/bin"

