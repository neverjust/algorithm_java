#!/bin/bash
# 打印操作
function printStep(){
  echo "### 执行操作【$1】###"
}

#判断 字符串参数1是否包含字符串参数2
function countainStr(){
  result=$(echo $1 | grep "${2}")
  if [[ "$result" != "" ]]
  then
    return 1
  else
    return 0
  fi
}
#ADD
echo -e "\n"
printStep "git add"
echo `git add .`

printStep "git status"
echo -e "\n"
statusResult=`git status`
echo $statusResult

# 如果没有文件修改
countainStr $statusResult "nothing to commit"
if [ $? == 1 ]
then
  echo "当前文件夹没有被【新建】或【修改】"
  exit
fi

# 提交内容为空
message="$1"
if [ "$message" = "" ]
then
  echo "请输入提交内容"
  read $message
fi

printStep "git commit -m ${message}"
echo `git commit -m ${message}`

printStep "git push"
pushResult=`git push -u origin master`

# 如果推送远程报错
countainStr $pushResult "fatal: "
if [ $? == 1 ]
then
  echo "推送远程小伙伴出错了，请联系开发小哥哥看看哪里出错了"
else
  echo "提交完成"
fi

