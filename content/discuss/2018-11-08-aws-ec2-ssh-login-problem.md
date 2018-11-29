---
title:  2018-11-08
---
来源：《Go 夜读》微信群

时间：2018-11-08

---

## aws ec2实例无法连接登录问题

在AWS上注册一个账号,然后创建一个免费试用期限的虚拟机,虚拟机成功创建并启动之后,我们通过SSH登录这台虚拟机出现无法连接的问题.

连接方式aws console页面会提示你如何连接,如下:

```bash
# 1. 创建虚拟机的时候会让你生成一个pem秘钥文件作为登录认证,把这个xxx.pem文件放到自己的机器上并设置400权限
chmod 400 wpc-secret.pem

# 2. 登录(注意我的是aws虚拟机是centos系统需要ec2-user用户登录,加-v参数显示登录过程详细信息)
ssh -i "wpc-secret.pem" ec2-user@ec2-18-224-136-148.us-east-2.compute.amazonaws.com -v

# 3. 我们会发现无法登录,超时登录后断开了,是因为默认ec2上面创建的虚拟机在默认情况下，默认安全组不允许传入SSH流量。所以你需要在ec2 console中的左侧菜单栏找到安全组然后添加SSH类型的入站和出站规则即可
略...

# 4. 重新尝试登录,一般情况先就成功登录上了
ssh -i "wpc-secret.pem" ec2-user@ec2-18-224-136-148.us-east-2.compute.amazonaws.com -v
...
```
