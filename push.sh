MESSAGE="$1"
if [ -z "${MESSAGE}" ]; then
echo "please input message "
exit 1
fi


set -x
time=`git status`
time=`git add .`
time=`git commit -m "${MESSAGE}"`
git push origin yun
set +x
exit 0
