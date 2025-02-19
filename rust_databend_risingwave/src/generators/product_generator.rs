use crate::models::product::Product;
use fake::{Fake};

use fake::faker::{
    barcode::zh_cn::Isbn,
    company::en::CompanyName,
    lorem::en::Paragraph,
    address::en::StateName,
};
use rand::Rng;

// 生成产品数据
pub fn generate_product_data() -> Product {
    // 生成产品名称
    let name: String = CompanyName().fake();
    // 生成随机数生成器
    let mut rng = rand::thread_rng();
    // 生成产品价格
    let price: f64 = rng.gen_range(10.0..1000.0);
    // 生成产品库存数量
    let stock_quantity: i32 = rng.gen_range(0..1000);

    // 创建产品
    let mut product = Product::new(name, price, stock_quantity);

    // 生成产品描述
    product.description = Some(Paragraph(1..3).fake());
    // 生成产品SKU
    product.sku = Some(Isbn().fake());
    product.category = Some(StateName().fake());
    product.brand = Some(CompanyName().fake());
    // 生成产品品牌
    product.weight = Some((rng.gen_range(0.0..100.0) * 100.0_f64).round() / 100.0);
    // 生成产品重量
    product.dimensions = Some(format!(
    // 生成产品尺寸
        "{}x{}x{}",
        rng.gen_range(10..100),
        rng.gen_range(10..100),
        rng.gen_range(10..100)
    ));
    product.color = Some("黑色".to_string());
    // 生成产品颜色
    product.material = Some("塑料".to_string());
    // 生成产品材质
    product.manufacturer = Some(CompanyName().fake());
    // 生成产品制造商
    product.supplier_id = Some(rng.gen_range(1..100));
    // 生成产品供应商ID
    product.min_stock_level = Some(rng.gen_range(1..50));
    // 生成产品最低库存水平
    product.max_stock_level = Some(rng.gen_range(100..200));
    // 生成产品最高库存水平
    product.reorder_quantity = Some(rng.gen_range(1..100));
    // 生成产品补货数量
    product.tax_rate = Some(0.0);
    // 生成产品税率
    product.discount_percentage = Some(0.0);
    // 生成产品折扣百分比
    product.rating = Some((rng.gen_range(0.0..5.0_f64) * 100.0).round() / 100.0);
    // 生成产品评分

    product
    // 返回产品
}