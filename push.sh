MESSAGE="$1"
if [ -z "${MESSAGE}" ]; then
echo "please input message "
exit 1
fi


set -x
git status
res=$?

   if [ $res -ne 0 ]; then
     echo "Failed to generate ${MESSAGE} config material..."
     exit 1
   fi


git add .
res=$?

    if [ $res -ne 0 ]; then
        echo "Failed to generate ${MESSAGE} config material..."
        exit 1
    fi


git commit -m "${MESSAGE}"
res=$?

    if [ $res -ne 0 ]; then
        echo "Failed to generate ${MESSAGE} config material..."
        exit 1
    fi


git push origin yun
res=$?

    if [ $res -ne 0 ]; then
        echo "Failed to generate ${MESSAGE} config material..."
        exit 1
    fi
set +x
exit 0
