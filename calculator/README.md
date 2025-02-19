计算使用databend的成本月使用费和年使用费 参考该文档 @https://docs.databend.cn/guides/products/dc/pricing
假设现有databend数据存储为500gb,购买计算集群为 商业版 每秒使用, 然后注意不使用的时候databend会在5分钟后自动暂停, 
假设一个小时查询5次databend,每次查询api的后5分钟内不会再次查询. 
分别按照不同的计算集群大小列出天和月和年的计算费用(每秒使用方式计算),存储费用,云服务费用和api调用费用以及总和的天,月,年费用# big_dataworks_databend_risingwave



需求如下:  polardb mysql 源数据库 => dataworks 走 mysql cdc订阅源 => 统计sql 编写 存储报表数据到小表里 => metbase , dataese 直出报表 

背景: 
polardb mysql 源数据库
1. 最大单表1亿-10亿之间；
2. 订单表千万级；
3. 用户表百万级；

数据库容量在35gb左右,每个月3 4gb增长..

现在在我自己的账号做试验,预估下费用和使用效果.评估后再走生产.

问题:
1 dataworks需要手动编写flink代码吗? 还是手动写一些统计sql 拖拉ui即可.
2 做一次聚合查询,例如统计 每日订单中的女性用户下单某商品的总数 大概的出数据时间
3 买什么size的包年包月套餐可覆盖现有数据库的增长 small  medium large?




阿里云datax 传输服务
https://github.com/alibaba/DataX

star星星的问题




计算使用databend的成本月使用费和年使用费 参考该文档
@https://docs.databend.cn/guides/products/dc/pricing   
假设现有databend数据存储为500gb,购买计算集群为 商业版 每秒使用,
然后注意不是用的时候databend会在5中后自动暂停,
假设一个小时查询5次databend,然后再5分钟内不会再次查询.
分别按照不同的计算集群大小列出天和月和年的计算费用(每秒使用方式计算),存储费用,云服务费用和api调用费用.





使用rust 批量生成 电商数据表,用户表,商品product表和订单明细表 三个表 注意三个表的实体联系.去掉外键关联,每个表需要有20个有意义的字段...然后 rust 批量生成数据 然后不停的批量插入到三个表中. 大概要生成35gb左右的数据. 全部生成完毕后.    继续生成三个表的数据 但是生成频率修改为每半分钟生成 三条数据插入三个表中...........
