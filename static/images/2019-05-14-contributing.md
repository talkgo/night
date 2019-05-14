CMLiang:git命令使用
//克隆分支到你本地
git clone https://github.com/CMLiang/reading-go
//查看列出详细信息，在每一个名字后面列出其远程url
git remote -v
origin  https://github.com/CMLiang/reading-go (fetch)
origin  https://github.com/CMLiang/reading-go (push)
//添加正式仓库地址
git remote add upstream https://github.com/developer-learning/night-reading-go.git
//让你本地 master 分支保持最新：
git fetch upstream
$ git checkout master
$ git rebase upstream/master
//从 master 开分支
git checkout -b myfeature
git commit -a -m "message"
git push -f origin myfeature

// (注：不要带--hard)到上个版本
git reset commitId
// 暂存修改
git stash
// 强制push,远程的最新的一次commit被删除
git push --force
// 释放暂存的修改，开始修改代码
git stash pop
// 增加文件流程
git add . -> git commit -m "massage" -> git push
