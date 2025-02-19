use chrono::{NaiveDate, NaiveDateTime, Utc};
use serde::{Deserialize, Serialize};
use crate::schema::users;

#[derive(Debug, Clone, Insertable, Queryable, Serialize, Deserialize, sqlx::FromRow)]
pub struct User {
    pub id: i32,
    pub username: String,
    pub email: String,
    pub password_hash: String,
    pub gender: Option<String>,
    pub first_name: String,
    pub last_name: String,
    pub phone: Option<String>,
    pub address: Option<String>,
    pub city: Option<String>,
    pub country: Option<String>,
    pub postal_code: Option<String>,
    pub birth_date: Option<NaiveDate>,
    pub registration_date: NaiveDateTime,
    pub last_login: Option<NaiveDateTime>,
    pub is_active: bool,
    pub account_balance: Option<f64>,
    pub loyalty_points: Option<i32>,
    pub preferences: Option<String>,
    pub avatar_url: Option<String>,
    pub marketing_consent: bool,
    pub created_at: NaiveDateTime,
}

impl User {
    /// 创建新用户，只需要必填信息，其余字段后续可更新
    pub fn new(username: String, email: String, password_hash: String) -> Self {
        let now = Utc::now().naive_utc();
        Self {
            id: 0,
            username,
            email,
            password_hash,
            gender: None,
            first_name: String::new(),
            last_name: String::new(),
            phone: None,
            address: None,
            city: None,
            country: None,
            postal_code: None,
            birth_date: None,
            registration_date: now,
            last_login: None,
            is_active: true,
            account_balance: None,
            loyalty_points: None,
            preferences: None,
            avatar_url: None,
            marketing_consent: false,
            created_at: now,
        }
    }
}