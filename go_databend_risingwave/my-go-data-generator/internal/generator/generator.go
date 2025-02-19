package generator

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"gorm.io/gorm"
	"my-go-data-generator/internal/models"
)

var (
	// 记录数由 CalculateRecordCounts 动态计算得出
	numUsers    int
	numProducts int
	numOrders   int
)

var (
	genders        = []string{"男", "女", "其他"}
	occupations    = []string{"工程师", "医生", "教师", "艺术家", "律师"}
	maritalStatus  = []string{"未婚", "已婚", "离异"}
	educationList  = []string{"高中", "本科", "硕士", "博士"}
	productNames   = []string{"产品A", "产品B", "产品C", "产品D", "产品E"}
	categories     = []string{"电子产品", "家居用品", "服装", "运动器材", "食品"}
	paymentMethods = []string{"信用卡", "支付宝", "微信支付", "现金"}
	orderStatuses  = []string{"待付款", "已付款", "待发货", "已发货", "已完成", "已取消"}
	stockStatuses  = []string{"有货", "缺货", "预订"}
)

func init() {
	// 以目标数据量 50GB 动态计算记录数
	numUsers, numProducts, numOrders = CalculateRecordCounts(50) // 50GB
	log.Printf("目标数据量设置：用户=%d, 产品=%d, 订单=%d", numUsers, numProducts, numOrders)
}

// CalculateRecordCounts 根据目标 GB 数据量及各表的预估平均行大小计算记录数
// 假设比例：产品数量 = 用户数量/10，订单数量 = 用户数量*10
func CalculateRecordCounts(totalGB int) (int, int, int) {
	// 预估每行大小（单位：字节），需要根据实际字段长度和数据内容调试
	estimatedUserRowSize := 300.0    // 例如用户记录约300字节
	estimatedProductRowSize := 400.0 // 产品记录约400字节
	estimatedOrderRowSize := 500.0   // 订单记录约500字节

	// 总体数据量 = 用户记录总字节 + 产品记录总字节 + 订单记录总字节
	//               = u*estimatedUserRowSize + (u/10)*estimatedProductRowSize + (u*10)*estimatedOrderRowSize
	factor := estimatedUserRowSize + estimatedProductRowSize/10 + 10*estimatedOrderRowSize
	totalBytes := float64(totalGB) * 1024 * 1024 * 1024
	u := totalBytes / factor
	return int(u), int(u) / 10, int(u) * 10
}

var (
	mutexUsers    sync.Mutex
	mutexProducts sync.Mutex
	mutexOrders   sync.Mutex
)

// GenerateData 并发生成用户、产品和订单数据，使用批量插入和并发提高性能
// CSV 写入部分已暂时注释掉
func GenerateData(db *gorm.DB) error {
	rand.Seed(time.Now().UnixNano())
	batchSize := 1000

	/*
		// 以下CSV相关代码暂时注释掉
		// 针对每个表建立单独的 CSV 写入 channel 和 writer goroutine
		userCSVChan := make(chan []string, 1000)
		productCSVChan := make(chan []string, 1000)
		orderCSVChan := make(chan []string, 1000)
		var csvWg sync.WaitGroup
		csvWg.Add(3)
		go WriteCSVConcurrently("users.csv", userCSVChan, &csvWg)
		go WriteCSVConcurrently("products.csv", productCSVChan, &csvWg)
		go WriteCSVConcurrently("orders.csv", orderCSVChan, &csvWg)

		// 写入CSV文件的表头
		userCSVChan <- []string{"ID", "Username", "Gender", "Age", "Email", "Phone", "Address", "Nationality", "Occupation", "MaritalStatus", "Education", "Hobby", "Income", "RegistrationDate", "LastLogin", "LoyaltyPoints", "PreferredLanguage", "Currency", "Timezone", "Status", "CreatedAt", "UpdatedAt"}
		productCSVChan <- []string{"ID", "ProductName", "Category", "Description", "Price", "Stock", "SKU", "Manufacturer", "Weight", "Dimensions", "Color", "Material", "ReleaseDate", "WarrantyPeriod", "CountryOfOrigin", "Rating", "NumberOfReviews", "Discount", "StockStatus", "Supplier", "CreatedAt", "UpdatedAt"}
		orderCSVChan <- []string{"ID", "OrderNumber", "UserID", "ProductID", "OrderDate", "Quantity", "TotalAmount", "PaymentMethod", "ShippingAddress", "BillingAddress", "OrderStatus", "DiscountAmount", "TaxAmount", "ShippingCost", "TrackingNumber", "DeliveryDate", "ReturnStatus", "CustomerNote", "InternalNote", "IsGift", "GiftMessage", "ExtraInfo", "CreatedAt", "UpdatedAt"}
	*/

	// 使用并发批量插入，每个批次使用一个 goroutine。限制并发数防止过多 goroutine
	maxWorkers := runtime.NumCPU() * 2
	sem := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup

	// 生成用户数据
	var allUsers []models.User
	log.Println("开始生成用户数据...")
	for i := 0; i < numUsers; i += batchSize {
		wg.Add(1)
		sem <- struct{}{}
		go func(start int) {
			defer wg.Done()
			var users []models.User
			for j := 0; j < batchSize && (start+j) < numUsers; j++ {
				index := start + j + 1 // 保证唯一性
				now := time.Now()
				user := models.User{
					Username:          fmt.Sprintf("用户%d", index),
					Gender:            genders[rand.Intn(len(genders))],
					Age:               rand.Intn(63) + 18,
					Email:             fmt.Sprintf("user%d@example.com", index),
					Phone:             fmt.Sprintf("138%08d", index),
					Address:           fmt.Sprintf("地址%d", index),
					Nationality:       "中国",
					Occupation:        occupations[rand.Intn(len(occupations))],
					MaritalStatus:     maritalStatus[rand.Intn(len(maritalStatus))],
					Education:         educationList[rand.Intn(len(educationList))],
					Hobby:             "阅读,旅行",
					Income:            rand.Float64()*10000 + 3000,
					RegistrationDate:  now.Add(-time.Hour * time.Duration(rand.Intn(10000))),
					LastLogin:         now.Add(-time.Minute * time.Duration(rand.Intn(10000))),
					LoyaltyPoints:     rand.Intn(1000),
					PreferredLanguage: "中文",
					Currency:          "CNY",
					Timezone:          "CST",
					Status:            "活跃",
					CreatedAt:         now,
					UpdatedAt:         now,
				}
				users = append(users, user)
				/*
					// CSV写入相关代码已暂时注释掉
					csvRecord := []string{
						strconv.Itoa(int(user.ID)),
						user.Username,
						user.Gender,
						strconv.Itoa(user.Age),
						user.Email,
						user.Phone,
						user.Address,
						user.Nationality,
						user.Occupation,
						user.MaritalStatus,
						user.Education,
						user.Hobby,
						fmt.Sprintf("%.2f", user.Income),
						user.RegistrationDate.Format(time.RFC3339),
						user.LastLogin.Format(time.RFC3339),
						strconv.Itoa(user.LoyaltyPoints),
						user.PreferredLanguage,
						user.Currency,
						user.Timezone,
						user.Status,
						user.CreatedAt.Format(time.RFC3339),
						user.UpdatedAt.Format(time.RFC3339),
					}
					userCSVChan <- csvRecord
				*/
			}
			if err := db.Create(&users).Error; err != nil {
				log.Printf("批量插入用户数据失败: %v", err)
			}
			mutexUsers.Lock()
			allUsers = append(allUsers, users...)
			mutexUsers.Unlock()
			log.Printf("已插入用户数据：%d/%d", len(allUsers), numUsers)
			<-sem
		}(i)
	}
	wg.Wait()
	log.Println("用户数据生成完毕.")

	// 生成产品数据
	var allProducts []models.Product
	log.Println("开始生成产品数据...")
	for i := 0; i < numProducts; i += batchSize {
		wg.Add(1)
		sem <- struct{}{}
		go func(start int) {
			defer wg.Done()
			var products []models.Product
			for j := 0; j < batchSize && (start+j) < numProducts; j++ {
				index := start + j + 1
				now := time.Now()
				product := models.Product{
					ProductName:     productNames[rand.Intn(len(productNames))] + fmt.Sprintf(" %d", index),
					Category:        categories[rand.Intn(len(categories))],
					Description:     fmt.Sprintf("这是%s的描述", productNames[rand.Intn(len(productNames))]),
					Price:           rand.Float64() * 1000,
					Stock:           rand.Intn(5000),
					SKU:             fmt.Sprintf("SKU%06d", index),
					Manufacturer:    fmt.Sprintf("制造商%d", rand.Intn(100)),
					Weight:          rand.Float64() * 10,
					Dimensions:      fmt.Sprintf("%dx%dx%d", rand.Intn(100), rand.Intn(100), rand.Intn(100)),
					Color:           []string{"红", "蓝", "绿", "黑", "白"}[rand.Intn(5)],
					Material:        "塑料",
					ReleaseDate:     now.AddDate(-rand.Intn(10), 0, 0),
					WarrantyPeriod:  fmt.Sprintf("%d个月", rand.Intn(24)+1),
					CountryOfOrigin: "中国",
					Rating:          rand.Float64() * 5,
					NumberOfReviews: rand.Intn(1000),
					Discount:        rand.Float64() * 0.5,
					StockStatus:     stockStatuses[rand.Intn(len(stockStatuses))],
					Supplier:        fmt.Sprintf("供应商%d", rand.Intn(50)),
					CreatedAt:       now,
					UpdatedAt:       now,
				}
				products = append(products, product)
				/*
					// CSV写入相关代码已暂时注释掉
					csvRecord := []string{
						strconv.Itoa(int(product.ID)),
						product.ProductName,
						product.Category,
						product.Description,
						fmt.Sprintf("%.2f", product.Price),
						strconv.Itoa(product.Stock),
						product.SKU,
						product.Manufacturer,
						fmt.Sprintf("%.2f", product.Weight),
						product.Dimensions,
						product.Color,
						product.Material,
						product.ReleaseDate.Format(time.RFC3339),
						product.WarrantyPeriod,
						product.CountryOfOrigin,
						fmt.Sprintf("%.2f", product.Rating),
						strconv.Itoa(product.NumberOfReviews),
						fmt.Sprintf("%.2f", product.Discount),
						product.StockStatus,
						product.Supplier,
						product.CreatedAt.Format(time.RFC3339),
						product.UpdatedAt.Format(time.RFC3339),
					}
					productCSVChan <- csvRecord
				*/
			}
			if err := db.Create(&products).Error; err != nil {
				log.Printf("批量插入产品数据失败: %v", err)
			}
			mutexProducts.Lock()
			allProducts = append(allProducts, products...)
			mutexProducts.Unlock()
			log.Printf("已插入产品数据：%d/%d", len(allProducts), numProducts)
			<-sem
		}(i)
	}
	wg.Wait()
	log.Println("产品数据生成完毕.")

	// 生成订单数据
	log.Println("开始生成订单数据...")
	orderBatchSize := 1000
	var countOrders int
	for i := 0; i < numOrders; i += orderBatchSize {
		wg.Add(1)
		sem <- struct{}{}
		go func(start int) {
			defer wg.Done()
			var orders []models.Order
			for j := 0; j < orderBatchSize && (start+j) < numOrders; j++ {
				now := time.Now()
				// 随机选择已生成的用户和产品作为关联数据
				var user models.User
				var product models.Product
				mutexUsers.Lock()
				if len(allUsers) > 0 {
					user = allUsers[rand.Intn(len(allUsers))]
				}
				mutexUsers.Unlock()
				mutexProducts.Lock()
				if len(allProducts) > 0 {
					product = allProducts[rand.Intn(len(allProducts))]
				}
				mutexProducts.Unlock()

				order := models.Order{
					OrderNumber:     fmt.Sprintf("ORD%010d", start+j+1),
					UserID:          user.ID,
					ProductID:       product.ID,
					OrderDate:       now.Add(-time.Duration(rand.Intn(1000)) * time.Minute),
					Quantity:        rand.Intn(10) + 1,
					TotalAmount:     product.Price * float64(rand.Intn(10)+1),
					PaymentMethod:   paymentMethods[rand.Intn(len(paymentMethods))],
					ShippingAddress: fmt.Sprintf("收货地址%d", start+j+1),
					BillingAddress:  fmt.Sprintf("账单地址%d", start+j+1),
					OrderStatus:     orderStatuses[rand.Intn(len(orderStatuses))],
					DiscountAmount:  rand.Float64() * 50,
					TaxAmount:       rand.Float64() * 20,
					ShippingCost:    rand.Float64() * 10,
					TrackingNumber:  fmt.Sprintf("TRK%08d", rand.Intn(100000000)),
					DeliveryDate:    now.Add(time.Duration(rand.Intn(1000)) * time.Minute),
					ReturnStatus:    "无",
					CustomerNote:    "请尽快发货",
					InternalNote:    "内部备注信息",
					IsGift:          rand.Intn(2) == 0,
					GiftMessage:     "祝您购物愉快",
					ExtraInfo:       "额外信息",
					CreatedAt:       now,
					UpdatedAt:       now,
				}
				orders = append(orders, order)
				/*
					// CSV写入相关代码已暂时注释掉
					csvRecord := []string{
						strconv.Itoa(int(order.ID)),
						order.OrderNumber,
						strconv.Itoa(int(order.UserID)),
						strconv.Itoa(int(order.ProductID)),
						order.OrderDate.Format(time.RFC3339),
						strconv.Itoa(order.Quantity),
						fmt.Sprintf("%.2f", order.TotalAmount),
						order.PaymentMethod,
						order.ShippingAddress,
						order.BillingAddress,
						order.OrderStatus,
						fmt.Sprintf("%.2f", order.DiscountAmount),
						fmt.Sprintf("%.2f", order.TaxAmount),
						fmt.Sprintf("%.2f", order.ShippingCost),
						order.TrackingNumber,
						order.DeliveryDate.Format(time.RFC3339),
						order.ReturnStatus,
						order.CustomerNote,
						order.InternalNote,
						fmt.Sprintf("%t", order.IsGift),
						order.GiftMessage,
						order.ExtraInfo,
						order.CreatedAt.Format(time.RFC3339),
						order.UpdatedAt.Format(time.RFC3339),
					}
					orderCSVChan <- csvRecord
				*/
			}
			if err := db.Create(&orders).Error; err != nil {
				log.Printf("批量插入订单数据失败: %v", err)
			}
			mutexOrders.Lock()
			countOrders += len(orders)
			mutexOrders.Unlock()
			if countOrders%(orderBatchSize*10) == 0 {
				log.Printf("已插入订单数据：%d/%d", countOrders, numOrders)
			}
			<-sem
		}(i)
	}
	wg.Wait()
	log.Println("订单数据生成完毕.")

	/*
		// 关闭所有 CSV 通道，等待 CSV 写入 goroutine 完成
		close(userCSVChan)
		close(productCSVChan)
		close(orderCSVChan)
		csvWg.Wait()
		log.Println("CSV文件写入完毕.")
	*/

	return nil
}

// StartTimer 启动定时器，每30秒向三个表中分别插入一条新数据，并执行 JOIN 查询打印结果及当前运行时长
func StartTimer(db *gorm.DB, startTime time.Time) {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for range ticker.C {
			now := time.Now()
			// 插入一条用户数据，确保手机号唯一
			user := models.User{
				Username:          fmt.Sprintf("定时用户%d", now.UnixNano()),
				Gender:            genders[rand.Intn(len(genders))],
				Age:               rand.Intn(63) + 18,
				Email:             fmt.Sprintf("timed_user%d@example.com", now.UnixNano()),
				Phone:             fmt.Sprintf("139%08d", now.UnixNano()%100000000),
				Address:           "定时地址",
				Nationality:       "中国",
				Occupation:        occupations[rand.Intn(len(occupations))],
				MaritalStatus:     maritalStatus[rand.Intn(len(maritalStatus))],
				Education:         educationList[rand.Intn(len(educationList))],
				Hobby:             "运动,音乐",
				Income:            rand.Float64()*10000 + 3000,
				RegistrationDate:  now,
				LastLogin:         now,
				LoyaltyPoints:     rand.Intn(1000),
				PreferredLanguage: "中文",
				Currency:          "CNY",
				Timezone:          "CST",
				Status:            "活跃",
				CreatedAt:         now,
				UpdatedAt:         now,
			}
			if err := db.Create(&user).Error; err != nil {
				log.Printf("定时插入用户失败: %v", err)
				continue
			}

			// 插入一条产品数据
			product := models.Product{
				ProductName:     fmt.Sprintf("定时产品%d", now.UnixNano()),
				Category:        categories[rand.Intn(len(categories))],
				Description:     "定时生成的产品描述",
				Price:           rand.Float64() * 1000,
				Stock:           rand.Intn(5000),
				SKU:             fmt.Sprintf("TSKU%06d", rand.Intn(1000000)),
				Manufacturer:    "定时制造商",
				Weight:          rand.Float64() * 10,
				Dimensions:      fmt.Sprintf("%dx%dx%d", rand.Intn(100), rand.Intn(100), rand.Intn(100)),
				Color:           []string{"红", "蓝", "绿", "黑", "白"}[rand.Intn(5)],
				Material:        "定时材质",
				ReleaseDate:     now,
				WarrantyPeriod:  "12个月",
				CountryOfOrigin: "中国",
				Rating:          rand.Float64() * 5,
				NumberOfReviews: rand.Intn(1000),
				Discount:        rand.Float64() * 0.5,
				StockStatus:     stockStatuses[rand.Intn(len(stockStatuses))],
				Supplier:        "定时供应商",
				CreatedAt:       now,
				UpdatedAt:       now,
			}
			if err := db.Create(&product).Error; err != nil {
				log.Printf("定时插入产品失败: %v", err)
				continue
			}

			// 插入一条订单数据，关联上述用户与产品
			order := models.Order{
				OrderNumber:     fmt.Sprintf("TORD%v", now.UnixNano()),
				UserID:          user.ID,
				ProductID:       product.ID,
				OrderDate:       now,
				Quantity:        rand.Intn(10) + 1,
				TotalAmount:     product.Price * float64(rand.Intn(10)+1),
				PaymentMethod:   paymentMethods[rand.Intn(len(paymentMethods))],
				ShippingAddress: "定时收货地址",
				BillingAddress:  "定时账单地址",
				OrderStatus:     "待付款",
				DiscountAmount:  rand.Float64() * 50,
				TaxAmount:       rand.Float64() * 20,
				ShippingCost:    rand.Float64() * 10,
				TrackingNumber:  fmt.Sprintf("TTRK%v", rand.Intn(1000000)),
				DeliveryDate:    now.Add(24 * time.Hour),
				ReturnStatus:    "无",
				CustomerNote:    "定时订单, 请尽快处理",
				InternalNote:    "定时内部备注",
				IsGift:          false,
				GiftMessage:     "",
				ExtraInfo:       "定时额外信息",
				CreatedAt:       now,
				UpdatedAt:       now,
			}
			if err := db.Create(&order).Error; err != nil {
				log.Printf("定时插入订单失败: %v", err)
				continue
			}

			// 使用 JOIN 查询刚刚插入的定时订单数据
			var joinedResult struct {
				OrderNumber string
				Username    string
				ProductName string
				TotalAmount float64
			}
			err := db.Table("orders").
				Select("orders.order_number, users.username, products.product_name, orders.total_amount").
				Joins("JOIN users ON orders.user_id = users.id").
				Joins("JOIN products ON orders.product_id = products.id").
				Where("orders.id = ?", order.ID).
				First(&joinedResult).Error
			if err != nil {
				log.Printf("查询定时订单失败: %v", err)
				continue
			}
			elapsed := time.Since(startTime)
			log.Printf("定时任务插入并查询成功：订单号：%s，用户：%s，产品：%s，金额：%.2f，运行时间：%s",
				joinedResult.OrderNumber, joinedResult.Username, joinedResult.ProductName, joinedResult.TotalAmount, elapsed)
		}
	}()
}
