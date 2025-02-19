计算使用databend的成本月使用费和年使用费 参考该文档 @https://docs.databend.cn/guides/products/dc/pricing
假设现有databend数据存储为500gb,购买计算集群为 商业版 每秒使用, 然后注意不使用的时候databend会在5分钟后自动暂停, 
假设一个小时查询5次databend,每次查询api的后5分钟内不会再次查询. 
分别按照不同的计算集群大小列出天和月和年的计算费用(每秒使用方式计算),存储费用,云服务费用和api调用费用以及总和的天,月,年费用