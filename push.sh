MESSAGE="$1"
if [ -z "${MESSAGE}" ]; then
echo "please input message "
exit 1
fi


set -x
git status
res=$?
set +x
   if [ $res -ne 0 ]; then
     echo "Failed to generate ${MESSAGE} config material..."
     exit 1
   fi

set -x
git add .
res=$?
set +x
    if [ $res -ne 0 ]; then
        echo "Failed to generate ${MESSAGE} config material..."
        exit 1
    fi

set -x
git commit -m "${MESSAGE}"
res=$?
set +x
    if [ $res -ne 0 ]; then
        echo "Failed to generate ${MESSAGE} config material..."
        exit 1
    fi

set -x
git push origin yun
res=$?
set +x
    if [ $res -ne 0 ]; then
        echo "Failed to generate ${MESSAGE} config material..."
        exit 1
    fi

