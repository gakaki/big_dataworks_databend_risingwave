use chrono::{NaiveDateTime, Utc};
use serde::{Deserialize, Serialize};
use crate::schema::products;

#[derive(Debug, Clone, Insertable, Queryable, Serialize, Deserialize, sqlx::FromRow)]
#[diesel(table_name = products)]
pub struct Product {
    pub id: i32,
    pub name: String,
    pub description: Option<String>,
    pub sku: Option<String>,
    pub price: f64,
    pub stock_quantity: i32,
    pub category: Option<String>,
    pub brand: Option<String>,
    pub weight: Option<f64>,
    pub dimensions: Option<String>,
    pub color: Option<String>,
    pub material: Option<String>,
    pub manufacturer: Option<String>,
    pub supplier_id: Option<i32>,
    pub min_stock_level: Option<i32>,
    pub max_stock_level: Option<i32>,
    pub reorder_quantity: Option<i32>,
    pub is_active: bool,
    pub tax_rate: Option<f64>,
    pub discount_percentage: Option<f64>,
    pub rating: Option<f64>,
    pub created_at: NaiveDateTime,
}

impl Product {
    /// 创建新产品，仅需提供名称、价格和库存，其他字段后续可更新
    pub fn new(name: String, price: f64, stock_quantity: i32) -> Self {
        let now = Utc::now().naive_utc();
        Self {
            id: 0,
            name,
            description: None,
            sku: None,
            price,
            stock_quantity,
            category: None,
            brand: None,
            weight: None,
            dimensions: None,
            color: None,
            material: None,
            manufacturer: None,
            supplier_id: None,
            min_stock_level: None,
            max_stock_level: None,
            reorder_quantity: None,
            is_active: true,
            tax_rate: None,
            discount_percentage: None,
            rating: None,
            created_at: now,
        }
    }
}