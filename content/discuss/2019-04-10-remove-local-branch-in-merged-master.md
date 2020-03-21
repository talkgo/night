---
title: 2019-04-10 
date: 2019-04-10T00:00:00+08:00
---

来源：『Go 夜读』微信群

## 本地开发分支已经 merge 到 master 后，如何快捷的删除本地开发分支？

```
#!/bin/bash

function git_branch_cleanup() {
    for branch in `git branch --format='%(refname:short)'|grep -v '\*\|master'` ; do
        git checkout $branch
        check_results=`git fetch origin master && git rebase origin/master`
        result=$(echo $check_results | grep "up to date.")
        if [ "$result" == "" ];then
            echo "不包含 up to date. $check_results !\n"
        fi
    done
    git checkout master
    git branch --merged | grep -v '\*\|master' | xargs -n 1 git branch -d
}

git_branch_cleanup
```

你可以将以上代码创建到 `/usr/local/bin/gcup`，这样你就可以在项目中使用 `gcup` 命令了。
>注意：gcup 需要权限：`chmod +x gcup`
>如果之前没有清理无用分支，可能会有大量冲突需要处理。

其他办法：
我们也可以遍历本地所有分支，然后排除一些白名单的分支，然后再把其他的全部删掉。
>`git branch -d xxx` 能删掉的就都删掉了，如果你确定 xxx 分支没有用了，你也可以强制删除 `git branch -D xxx`

更简洁的代码：`git branch -d $(git branch -vv | grep ': gone\]' | awk '{print $1}')`