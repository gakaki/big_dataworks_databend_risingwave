## Use Rust to Generate 500GB data in mysql database
## promopt
使用中文回答.
使用rust生成postgresql数据库大概要生成50gb的数据库.
三个表 订单order 产品product 用户use表,注意每个表都要有外键id关联,但需要去除外键强制依赖..也就是逻辑依赖即可.
每个表需要生成有意义的字段20个以上.每个字段需要生成有意义的值.
每个表需要生成有意义的索引.
批量高效的插入数据到postgresql数据库中.注意插入的数据要有关联性可以用select join三个表被查询出来.
生成的数据也需要写入csv文件中,方便后续使用postgresql命令行工具进行数据导入.
插入一定量的数据之后要打印log显示当前生成和插入数据以及写入csv的进度播报..
统计并打印出来运行了多长时间

完成指定数据量的生成和插入数据库和csv之后.
定时每半分钟 继续插入一条数据到三个数据表中,插入的数据需要有关联性.
然后使用select join三个表查询刚才插入的那条数据.并且打印出来.
最后还要统计并打印出来运行了多长时间



用户表增加性别 
然后使用disel的migration生成表结构和 给出指定的命令行

cargo install diesel_cli --no-default-features --features  mysql

brew install cargo-binstall
cargo binstall diesel_cli

export DATABASE_URL=postgres://username:password@localhost/database_name
export DATABASE_URL=mysql://root:mypassword@127.0.0.1:3306/mydb

不得使用localhost 只能用127的ip链接

diesel setup
diesel migration run

cargo build --release
cargo run --release


sudo apt install -y libssl-dev libmysqlclient-dev libmysqlclient21 librust-mysqlclient-sys-dev


