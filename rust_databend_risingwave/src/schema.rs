// @generated automatically by Diesel CLI.

diesel::table! {
    orders (id) {
        id -> Integer,
        user_id -> Integer,
        product_id -> Integer,
        quantity -> Integer,
        order_date -> Nullable<Datetime>,
        #[max_length = 50]
        status -> Nullable<Varchar>,
        shipping_address -> Nullable<Text>,
        #[max_length = 50]
        payment_method -> Nullable<Varchar>,
        total_amount -> Decimal,
        discount -> Nullable<Decimal>,
        tax -> Nullable<Decimal>,
        created_at -> Nullable<Datetime>,
        updated_at -> Nullable<Datetime>,
        #[max_length = 100]
        tracking_number -> Nullable<Varchar>,
        delivery_date -> Nullable<Datetime>,
        customer_feedback -> Nullable<Text>,
        is_gift -> Nullable<Bool>,
        gift_message -> Nullable<Text>,
        #[max_length = 50]
        order_source -> Nullable<Varchar>,
        #[max_length = 10]
        currency -> Nullable<Varchar>,
        #[max_length = 50]
        order_type -> Nullable<Varchar>,
    }
}

diesel::table! {
    products (id) {
        id -> Integer,
        #[max_length = 255]
        name -> Varchar,
        description -> Nullable<Text>,
        #[max_length = 100]
        sku -> Nullable<Varchar>,
        price -> Decimal,
        stock_quantity -> Nullable<Integer>,
        #[max_length = 100]
        category -> Nullable<Varchar>,
        #[max_length = 100]
        brand -> Nullable<Varchar>,
        weight -> Nullable<Decimal>,
        #[max_length = 100]
        dimensions -> Nullable<Varchar>,
        #[max_length = 50]
        color -> Nullable<Varchar>,
        #[max_length = 100]
        material -> Nullable<Varchar>,
        #[max_length = 255]
        manufacturer -> Nullable<Varchar>,
        supplier_id -> Nullable<Integer>,
        min_stock_level -> Nullable<Integer>,
        max_stock_level -> Nullable<Integer>,
        reorder_quantity -> Nullable<Integer>,
        is_active -> Nullable<Bool>,
        tax_rate -> Nullable<Decimal>,
        discount_percentage -> Nullable<Decimal>,
        rating -> Nullable<Decimal>,
        created_at -> Nullable<Datetime>,
    }
}

diesel::table! {
    users (id) {
        id -> Integer,
        #[max_length = 255]
        username -> Varchar,
        #[max_length = 255]
        email -> Varchar,
        #[max_length = 255]
        password_hash -> Varchar,
        #[max_length = 20]
        gender -> Nullable<Varchar>,
        #[max_length = 255]
        first_name -> Nullable<Varchar>,
        #[max_length = 255]
        last_name -> Nullable<Varchar>,
        #[max_length = 20]
        phone -> Nullable<Varchar>,
        address -> Nullable<Text>,
        #[max_length = 100]
        city -> Nullable<Varchar>,
        #[max_length = 100]
        country -> Nullable<Varchar>,
        #[max_length = 20]
        postal_code -> Nullable<Varchar>,
        birth_date -> Nullable<Date>,
        registration_date -> Nullable<Datetime>,
        last_login -> Nullable<Datetime>,
        is_active -> Nullable<Bool>,
        account_balance -> Nullable<Decimal>,
        loyalty_points -> Nullable<Integer>,
        preferences -> Nullable<Text>,
        avatar_url -> Nullable<Text>,
        marketing_consent -> Nullable<Bool>,
        created_at -> Nullable<Datetime>,
    }
}

diesel::allow_tables_to_appear_in_same_query!(
    orders,
    products,
    users,
);
