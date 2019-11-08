### 介绍
	文本过滤+字符过滤+敏感词过滤+关键词过滤+脏词过滤+多字符串匹配+Aho-Corasick 算法+Trie树+DFA 
        支持所有的日志格式
### 配置
	etc/log.toml
	log_list = [] //日志文件列表
	fillter_list =[] //过滤字符串列表
	find_list  = [] //查找字符串列表 注意：find_list和fillter_list是互斥的
	ding_webhook_url //钉钉url
	server_log //日志
	tail_line = "10" //tail -n 参数(从第多少行展示开始监听)
