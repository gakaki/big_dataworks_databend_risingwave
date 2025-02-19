use diesel::prelude::*;
use diesel::mysql::MysqlConnection;
use diesel::connection::Connection;
use crate::models::user::User;
use crate::models::product::Product;
use crate::models::order::Order;
use crate::schema::{users, products, orders};
use std::error::Error;

pub struct Database {
    pub conn: MysqlConnection,
}

impl Database {
    pub fn new(database_url: &str) -> Result<Self, Box<dyn Error>> {
        let conn = MysqlConnection::establish(database_url)?;
        Ok(Self { conn })
    }

    pub fn insert_users(&mut self, users: &[User]) -> Result<(), Box<dyn Error>> {
        diesel::insert_into(users::table)
            .values(users.iter().cloned())
            .execute(&mut self.conn)?;
        Ok(())
    }
    

    // 批量插入产品数据
    pub fn insert_products(&mut self, products: &[Product]) -> Result<(), Box<dyn Error>> {
        diesel::insert_into(products::table)
            .values(products.iter().cloned())
            .execute(&mut self.conn)?;
        Ok(())
    }
    // 批量插入订单数据

    pub fn insert_orders(&mut self, orders: &[Order]) -> Result<(), Box<dyn Error>> {
        diesel::insert_into(orders::table)
            .values(orders.iter().cloned())
            .execute(&mut self.conn)?;
        Ok(())
    }
    /// 使用JOIN查询最新插入的数据
    pub fn query_joined_data(&self, order_id: i32) -> Result<(User, Product, Order), Box<dyn Error>> {
        use schema::orders::dsl::*;
        use schema::users::dsl::*;
        use schema::products::dsl::*;

        // let result = orders::table
        //     .inner_join(users::table)
        //     .inner_join(products::table)
        //     .filter(orders::id.eq(order_id))
        //     .first::<(Order, User, Product)>(&self.conn)?;

        let result = orders::table
            .inner_join(users::table)
            .inner_join(products::table)
            .select((orders::all_columns, users::all_columns, products::all_columns))
            .first::<(order::Order, User, product::Product)>(&mut self.conn)?;
        Ok((result.1, result.2, result.0))
    }
}