use crate::models::user::User;
use chrono::{NaiveDate};
use fake::faker::{
    address::zh_cn::{CityName, CountryName, SecondaryAddress, ZipCode},
    internet::zh_cn::{FreeEmail, Username},
    name::en::Name,
    phone_number::en::PhoneNumber,
};
use fake::Fake;
use rand::Rng;

pub fn generate_user() -> User {
    let username: String = Username().fake();
    let email: String = FreeEmail().fake();
    let password_hash = "dummy_hash".to_string(); // 实际项目中请使用安全的哈希算法

    let mut user = User::new(username, email, password_hash);

    let mut rng = rand::thread_rng();
    let genders = ["Male", "Female", "Other"];
    user.gender = Some(genders[rng.gen_range(0..genders.len())].to_string());
    user.first_name = Name().fake();
    user.last_name = Name().fake();
    user.phone = Some(PhoneNumber().fake());
    user.address = Some(SecondaryAddress().fake());
    user.city = Some(CityName().fake());
    user.country = Some(CountryName().fake());
    user.postal_code = Some(ZipCode().fake());
    user.birth_date = Some(
        NaiveDate::from_ymd_opt(
            1970 + rng.gen_range(0..30),
            rng.gen_range(1..13),
            rng.gen_range(1..29),
        )
        .unwrap(),
    );
    user.marketing_consent = rng.gen_bool(0.5);
    user.account_balance = Some(((rng.gen_range(0.0..10000.0) * 100.0) as f64).round() as f64 / 100.0);
    user.loyalty_points = Some(rng.gen_range(0..1000));
    user.preferences = Some("{\"theme\": \"dark\"}".to_string());
    user.avatar_url = Some("https://example.com/avatar.png".to_string());

    user
}