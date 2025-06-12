package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB // Biến global để chứa instance DB

// InitDB khởi tạo kết nối đến MySQL và trả về GORM DB instance
func InitDB(cfg *Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	// Cấu hình GORM logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Ngưỡng SQL chậm
			LogLevel:                  logger.Info, // Mức log (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,        // Bỏ qua lỗi 'record not found'
			Colorful:                  true,        // Output màu mè
		},
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Sử dụng tên bảng số ít, ví dụ: 'user' thay vì 'users'
		},
		Logger: newLogger,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic database object: %w", err)
	}

	// Cấu hình connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connection established successfully.")
	return DB, nil
}

// GetDB trả về instance DB đã được khởi tạo
// Bạn có thể sử dụng hàm này ở các nơi khác trong ứng dụng để lấy DB instance
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database not initialized. Call InitDB() first.")
	}
	return DB
}
