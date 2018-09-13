MESSAGE="$1"
if [ -z "${MESSAGE}" ]; then
echo "please input message "
exit 1
fi

git status
git add .
git commit -m "${MESSAGE}"
git push origin yun
