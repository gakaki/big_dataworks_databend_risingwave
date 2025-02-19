use chrono::{NaiveDateTime, Utc};
use serde::{Deserialize, Serialize};
use crate::schema::orders;

#[derive(Debug, Clone, Insertable, Queryable, Serialize, Deserialize, sqlx::FromRow)]
pub struct Order {
    pub id: i32,
    pub user_id: i32,
    pub product_id: i32,
    pub quantity: i32,
    pub order_date: NaiveDateTime,
    pub status: String,
    pub shipping_address: String,
    pub payment_method: String,
    pub total_amount: f64,
    pub discount: f64,
    pub tax: f64,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
    pub tracking_number: Option<String>,
    pub delivery_date: Option<NaiveDateTime>,
    pub customer_feedback: Option<String>,
    pub is_gift: bool,
    pub gift_message: Option<String>,
    pub order_source: Option<String>,
    pub currency: String,
    pub order_type: Option<String>,
}

impl Order {
    /// 创建新的Order实例，注意id、created_at和updated_at由数据库生成
    pub fn new(user_id: i32, product_id: i32, quantity: i32, status: String, shipping_address: String, payment_method: String, total_amount: f64) -> Self {
        let now = Utc::now().naive_utc();
        Order {
            id: 0,
            user_id,
            product_id,
            quantity,
            order_date: now,
            status,
            shipping_address,
            payment_method,
            total_amount,
            discount: 0.0,
            tax: 0.0,
            created_at: now,
            updated_at: now,
            tracking_number: None,
            delivery_date: None,
            customer_feedback: None,
            is_gift: false,
            gift_message: None,
            order_source: None,
            currency: "USD".to_string(),
            order_type: None,
        }
    }
}