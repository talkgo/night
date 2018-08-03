# 如何参与贡献？

>欢迎大家参与讨论，更欢迎大家多多的回馈社区，来开始共享吧！

## 贡献流程

### 第一步：Fork [https://github.com/developer-learning/night-reading-go](https://github.com/developer-learning/night-reading-go) 项目

1. 访问 [https://github.com/developer-learning/night-reading-go](https://github.com/developer-learning/night-reading-go)；
2. 点击 Fork 按钮（顶部右侧），建立基于此的分支；

### 第二步：克隆分支到你本地

```sh
# Define a local working directory:
$ working_dir=/.../src/github.com/developer-learning
$ user={your github profile name}
$ mkdir -p $working_dir
$ cd $working_dir
$ git clone https://github.com/$user/night-reading-go.git
$ cd $working_dir/night-reading-go
$ git remote -v
origin	https://github.com/$user/night-reading-go.git (fetch)
origin	https://github.com/$user/night-reading-go.git (push)

$ git remote add upstream https://github.com/developer-learning/night-reading-go.git
$ git remote -v
origin	https://github.com/$user/night-reading-go.git (fetch)
origin	https://github.com/$user/night-reading-go.git (push)
upstream	https://github.com/developer-learning/night-reading-go.git (fetch)
upstream	https://github.com/developer-learning/night-reading-go.git (push)

# Never push to upstream master since you do not have write access.
$ git remote set-url --push upstream no_push
$ git remote -v
origin	https://github.com/$user/night-reading-go.git (fetch)
origin	https://github.com/$user/night-reading-go.git (push)
upstream	https://github.com/developer-learning/night-reading-go.git (fetch)
upstream	no_push (push)
```

### 第三步：分支

让你本地 master 分支保持最新：

```sh
$ cd $working_dir/night-reading-go
$ git fetch upstream
$ git checkout master
$ git rebase upstream/master
```

从 master 开分支：

```sh
$ git checkout -b myfeature
```

### 第四步：开发

#### 编辑代码

你现在能在 `myfeature` 分支上编辑代码/文档了。

请按照以下一些格式编写：

文件命名：`2018-08-03-contributing.md`；
文件的图片：放到 `images` 下；
文件内容：需要明确标题、日期等基本信息；

### 第五步：保持分支同步

```sh
# While on your myfeature branch.
$ git fetch upstream
$ git rebase upstream/master
```

### 第六步：提交

提交你的修改：

```sh
$ git commit
```

### 第七步：推送

准备好审核：

```sh
git push -f origin myfeature
```

### 第八步：创建一个 pull request

1. 访问你 fork 的 [https://github.com/$user/night-reading-go](https://github.com/$user/night-reading-go) (替换 $user)；
2. 点击 myfeature 分支旁边的 Compare & pull request 按钮；

### 第九步：获取代码审核

一旦你的 Pull Request 被打开，它将被分配给审核者。
这些审核人员将进行彻底的代码审查，寻找正确性，错误，改进机会，文档和评论以及样式。

