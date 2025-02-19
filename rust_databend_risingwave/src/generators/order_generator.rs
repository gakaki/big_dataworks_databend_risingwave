use rand::Rng;
use chrono::{Utc};
use crate::models::order::Order;

pub fn generate_order(user_id: i32, product_id: i32) -> Order {
    let mut rng = rand::thread_rng();
    
    Order {
        id: rng.gen_range(1..1_000_000),
        user_id,
        product_id,
        quantity: rng.gen_range(1..10),
        order_date: Utc::now().naive_utc(),
        status: if rng.gen_bool(0.8) { "completed".to_string() } else { "pending".to_string() },
        shipping_address: format!("{} {} St.", rng.gen_range(1..1000), rng.gen_range(1..100)),
        payment_method: if rng.gen_bool(0.5) { "credit_card".to_string() } else { "paypal".to_string() },
        total_amount: rng.gen_range(10.0..500.0),
        discount: rng.gen_range(0.0..50.0),
        tax: rng.gen_range(0.0..20.0),
        created_at: Utc::now().naive_utc(),
        updated_at: Utc::now().naive_utc(),
        tracking_number: Some(format!("TRACK-{}", rng.gen_range(1000..9999))),
        delivery_date: Some((Utc::now() + chrono::Duration::days(rng.gen_range(1..10))).naive_utc()),
        customer_feedback: if rng.gen_bool(0.3) { Some("Great service!".to_string()) } else { None },
        is_gift: rng.gen_bool(0.2),
        gift_message: if rng.gen_bool(0.2) { Some("Happy Birthday!".to_string()) } else { None },
        order_source: Some(if rng.gen_bool(0.5) { "website".to_string() } else { "mobile_app".to_string() }),
        currency: "USD".to_string(),
        order_type: Some(if rng.gen_bool(0.5) { "regular".to_string() } else { "express".to_string() }),
    }
}

