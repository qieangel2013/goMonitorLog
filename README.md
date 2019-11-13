### 介绍
	基于Aho-Corasick 算法+Trie树+DFA 来实现实时匹配，可以实时对文本过滤+字符过滤+敏感词过滤+关键词过滤+脏词过滤+多字符串匹配等处理
	支持所有的文本文件类型（包括所有日志）
### 配置
	etc/log.toml
	log_list = [] //日志文件列表
	fillter_list =[] //过滤字符串列表
	find_list  = [] //查找字符串列表，注意 find_list和fillter_list是互斥的
	ding_webhook_url //钉钉url
	server_log //日志
	tail_line = "10" //tail -n 参数(从第多少行展示开始监听)
### 如果你对我的辛勤劳动给予肯定，请给我捐赠，你的捐赠是我最大的动力
![](https://github.com/qieangel2013/zys/blob/master/public/images/pw.jpg)
![](https://github.com/qieangel2013/zys/blob/master/public/images/pay.png)
[项目捐赠列表](https://github.com/qieangel2013/zys/wiki/%E9%A1%B9%E7%9B%AE%E6%8D%90%E8%B5%A0)
