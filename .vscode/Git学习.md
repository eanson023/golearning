#Git

	工作区（Working Directory） 添加，编辑修改文件等动作
	暂存区 暂存已经修改的文件最后统一提交到git仓库中
	Git Repository（Git仓库）最终确认的文件保存到仓库，称为一个新的版本，并且对他人可见

常用命令
	首先最基础的是需要配置用户信息
		$ git config --global user.name "lanya"
		$ git config --global user.email shenglanya@corp.netease.com
		该设置在github仓库主页显示睡提交了该文件
		git init 该命令将创建一个名为 .git 的子目录，这个子目录含有你在初始化的Git仓库中所有的必须文件，这些文件是Git仓库的骨干。但是，在这个时候，我们仅仅是做了一个初始化的操作，你的项目里的文件还没有被跟踪。
		
	git log                         查看日志
	
	git reflog                     查看提交记录
	
	git status                    查看本地仓库当前的状态
	
	git reset --hard            版本回退（删除本地提交）

###提交文件
	git add xxx工作区提交到暂存区
	git status查看状态
	git commit -m "提交描述" 暂存区提交到仓库区
###修改文件
	与提交一样
###删除wenj
	1.删除本地仓库文件
	2.删除暂存区的文件git rm xxx
	3.提交 git commit -m '描述'
	
##Git远程仓库
	作用：备份，实现代码共享集中化管理
	git push
	克隆：
		git clone 仓库地址
		目的：将远程仓库(github对应的项目)下载到本地
	创建新仓库
		git init
		git add xxx
		git commit -m ''
		git remote add origin https://github.com/eanson1999/仓库名.get
		git push -u origin master
		事先要有远程仓库
##使用Github
	目的：借助github托管项目代码
	基本概念：
		仓库（repository）
		仓库用来存放项目代码，每个项目对应一个仓库，多个开元项目有多个仓库
		
		收藏(Star)
		仓库主页star按钮，意思为收藏项目的人数
		收藏项目方便下次查看
		
		复制克隆项目（Fork）
			该fork项目是独立存在的
			
		发起请求（Pull request）
			fork的项目发起请求 让原作者更新
		
		关注（Watch）
			看到个项目 ，关注后有任何更新 会收到通知
		
		事务卡片（Issue）
		发现代码BUG,但是目前没有成型代码，需要讨论


​		
​		Github主页
​		
		仓库主页
		
		个人主页
		
		创建仓库/创建项目

##Github Pages搭建网站
	https://用户名.github.io
	1)创建个人站点 -> 新建仓库（注：仓库必须是[用户名.github.io]）
	2）在仓库下创建index.html即可
	
	即可通过https://eanson1999.github.io来访问
	经支持静态网页
	
	项目站点:
		https://用户名.github.io/仓库名
		1）进入项目主页 点击setting 点击[launch automic page generate]
