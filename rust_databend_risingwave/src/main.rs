use chrono::Utc;
use csv::Writer;
use db::Database;
use indicatif::{ProgressBar, ProgressStyle};
use std::env;
use std::fs;
use std::time::{Duration, Instant};
use tokio::time;

mod db;
mod generators;
mod models;
#[macro_use]
pub mod schema;
const TARGET_SIZE: u64 = 50 * 1024 * 1024 * 1024; // 50GB
const BATCH_SIZE: usize = 1000; // 每批写入1000条记录

/// 写入 CSV 表头
fn write_csv_headers(
    user_writer: &mut csv::Writer<std::fs::File>,
    product_writer: &mut csv::Writer<std::fs::File>,
    order_writer: &mut csv::Writer<std::fs::File>,
) -> Result<(), Box<dyn std::error::Error>> {
    user_writer.write_record(&[
        "id", "username", "email", "password_hash", "gender", "first_name", "last_name", "phone",
        "address", "city", "country", "postal_code", "birth_date", "registration_date", "last_login",
        "is_active", "account_balance", "loyalty_points", "preferences", "avatar_url", "marketing_consent", "created_at",
    ])?;
    product_writer.write_record(&[
        "id", "name", "description", "sku", "price", "stock_quantity", "category", "brand", "weight",
        "dimensions", "color", "material", "manufacturer", "supplier_id", "min_stock_level", "max_stock_level",
        "reorder_quantity", "is_active", "tax_rate", "discount_percentage", "rating", "created_at",
    ])?;
    order_writer.write_record(&[
        "id", "user_id", "product_id", "quantity", "order_date", "status", "shipping_address",
        "payment_method", "total_amount", "discount", "tax", "created_at", "updated_at", "tracking_number",
        "delivery_date", "customer_feedback", "is_gift", "gift_message", "order_source", "currency", "order_type",
    ])?;
    user_writer.flush()?;
    product_writer.flush()?;
    order_writer.flush()?;
    Ok(())
}

/// 持续每半分钟插入一条数据并写入 CSV
async fn continuous_insertion(
    db: &Database,
    user_writer: &mut Writer<std::fs::File>,
    product_writer: &mut Writer<std::fs::File>,
    order_writer: &mut Writer<std::fs::File>,
) -> Result<(), Box<dyn std::error::Error>> {
    let mut interval = tokio::time::interval(Duration::from_secs(30));
    
    loop { 
        interval.tick().await;
        let insert_start = Instant::now();

        let user = generators::user_generator::generate_user();
        let product = generators::product_generator::generate_product_data();
        let order = generators::order_generator::generate_order(user.id, product.id);

        // 单条数据插入
        db.insert_users(&[user.clone()])?;
        db.insert_products(&[product.clone()])?;
        db.insert_orders(&[order.clone()])?;

        // CSV 写入部分保持不变
        user_writer.serialize(&user)?;
        product_writer.serialize(&product)?;
        order_writer.serialize(&order)?;
        user_writer.flush()?;
        product_writer.flush()?;
        order_writer.flush()?;

        println!("New records inserted at: {}", Utc::now());
        println!("Insertion took: {:?}", insert_start.elapsed());

        // 使用 JOIN 查询验证数据
        let (inserted_user, inserted_product, inserted_order) = db.query_joined_data(order.id)?;
        println!("Joined data: User={:?}, Product={:?}, Order={:?}", 
                 inserted_user, inserted_product, inserted_order);
    }
}
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // 加载.env文件
    dotenv::dotenv().ok();
    
    env_logger::init();

    let start = Instant::now();
    let database_url = env::var("DATABASE_URL")?;
    let db = Database::new(&database_url)?;


    // 创建 CSV 写入器
    let mut user_writer = Writer::from_path("users.csv")?;
    let mut product_writer = Writer::from_path("products.csv")?;
    let mut order_writer = Writer::from_path("orders.csv")?;

    // 写入 CSV 表头
    write_csv_headers(&mut user_writer, &mut product_writer, &mut order_writer)?;

    // 批量数据缓冲
    let mut users_batch = Vec::with_capacity(BATCH_SIZE);
    let mut products_batch = Vec::with_capacity(BATCH_SIZE);
    let mut orders_batch = Vec::with_capacity(BATCH_SIZE);

    // Initial bulk data generation
    let total_records = 1_000_000; // 根据需求调整数据量
    let pb = ProgressBar::new(total_records as u64);
    pb.set_style(
        ProgressStyle::default_bar()
            .template("[{elapsed_precise}] {bar:40.cyan/blue} {pos:>7}/{len:7} {msg}")
            .unwrap(),
    );

    for i in 0..total_records {
        let user = generators::user_generator::generate_user();
        let product = generators::product_generator::generate_product_data();
        let order = generators::order_generator::generate_order(user.id, product.id);

        // 将数据保存到缓冲区
        users_batch.push(user.clone());
        products_batch.push(product.clone());
        orders_batch.push(order.clone());

        // 同时写入 CSV 文件（逐条写入）
        user_writer.serialize(&user)?;
        product_writer.serialize(&product)?;
        order_writer.serialize(&order)?;
        pb.inc(1);

        // 每批达到 BATCH_SIZE 后一次性写入 MySQL 数据库
        if (i + 1) % BATCH_SIZE == 0 {
            // 使用 diesel 批量插入
            db.insert_users(&users_batch)?;
            db.insert_products(&products_batch)?;
            db.insert_orders(&orders_batch)?;
            
            users_batch.clear();
            products_batch.clear();
            orders_batch.clear();

            pb.set_message(format!("Generated {} records", i + 1));
            user_writer.flush()?;
            product_writer.flush()?;
            order_writer.flush()?;

            // 计算 CSV 文件大小，总和作为数据体积的近似值
            let user_size = fs::metadata("users.csv")?.len();
            let product_size = fs::metadata("products.csv")?.len();
            let order_size = fs::metadata("orders.csv")?.len();
            let total_size = user_size + product_size + order_size;
            println!(
                "当前CSV数据大小: {:.2}GB",
                total_size as f64 / (1024.0 * 1024.0 * 1024.0)
            );
            if total_size >= TARGET_SIZE {
                println!("达到50GB数据容量，停止初始批量生成。");
                break;
            }
        }
    }

    pb.finish_with_message("Bulk data generation completed");
    println!("Initial data generation took: {:?}", start.elapsed());

    // 每半分钟继续插入数据并写入 CSV（独立函数调用）
    continuous_insertion(&db, &mut user_writer, &mut product_writer, &mut order_writer).await?;

    Ok(())
}
