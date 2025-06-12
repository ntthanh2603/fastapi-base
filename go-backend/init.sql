
--  table user
CREATE TABLE drunk_user (
    account_id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Account ID',
    username VARCHAR(255) NOT NULL COMMENT 'Username',
    password VARCHAR(255) NOT NULL COMMENT 'Password',
    email VARCHAR(255) NOT NULL COMMENT 'Email address',
    status TINYINT NOT NULL DEFAULT 1 COMMENT 'Status: 1-active, 0-disabled',
    language VARCHAR(50) DEFAULT 'en' COMMENT 'Language setting',
    is_vip BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'Flag indicating if the user is a VIP',
    last_login_time BIGINT DEFAULT 0 COMMENT 'Last login time',
    create_time BIGINT NOT NULL DEFAULT (UNIX_TIMESTAMP()) COMMENT 'Creation time',
    update_time BIGINT NOT NULL DEFAULT (UNIX_TIMESTAMP()) COMMENT 'Update time',
    is_deleted TINYINT NOT NULL DEFAULT 0 COMMENT 'Soft delete flag: 0=active, 1=deleted'
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_unicode_ci;

-- table venues (quán nhậu, nhà hàng...)
DROP TABLE IF EXISTS `venues`;

CREATE TABLE `venues` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for the venue',
    `name` VARCHAR(255) NOT NULL COMMENT 'Name of the venue',
    `address` TEXT NULL COMMENT 'Address of the venue',
    `owner_user_id` BIGINT UNSIGNED NULL COMMENT 'ID of the user who owns/manages this venue, FK to users(id)',
    `location` POINT NULL COMMENT 'Geographical location of the venue (longitude, latitude)',
    `rating` DECIMAL(3, 2) NULL COMMENT 'Average rating score, e.g., 4.50',
    `review_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Total number of reviews counted for the average rating',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_venue_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Information about venues/pubs';

-- table voucher_templates

DROP TABLE IF EXISTS `voucher_templates`;

CREATE TABLE `voucher_templates` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for the voucher template',
    `name` VARCHAR(255) NOT NULL COMMENT 'Voucher template name (e.g., "VIP Discount 200k")',
    `description` TEXT NULL COMMENT 'Optional description of the voucher',

    `voucher_type` ENUM('FIXED_AMOUNT_DISCOUNT', 'PERCENTAGE_DISCOUNT') NOT NULL DEFAULT 'FIXED_AMOUNT_DISCOUNT' COMMENT 'Type of voucher',
    `min_spend_amount` DECIMAL(10,2) NOT NULL COMMENT 'Minimum spend required to apply the voucher',
    `discount_amount` DECIMAL(10,2) NOT NULL COMMENT 'Amount to be discounted',

    `issuer_type` ENUM('PLATFORM', 'VENUE') NOT NULL DEFAULT 'PLATFORM' COMMENT 'Indicates who issues the voucher',
    `issuer_venue_id` INT UNSIGNED NULL COMMENT 'ID of the venue if issuer_type is VENUE, foreign key to venues(id)',

    `valid_from` DATE NOT NULL COMMENT 'Date from which the voucher template is valid',
    `valid_until` DATE NOT NULL COMMENT 'Date until which the voucher template is valid',

    `usage_scope` ENUM('ENTIRE_VENUE_ORDER', 'SPECIFIC_ITEMS') NOT NULL DEFAULT 'ENTIRE_VENUE_ORDER' COMMENT 'Scope of usage for the voucher',

    `target_user_type` ENUM('ALL', 'NEW', 'VIP', 'DORMANT') NOT NULL DEFAULT 'VIP' COMMENT 'Target user type for this voucher',

    `max_total_issues` INT UNSIGNED NULL COMMENT 'Optional: Maximum total number of vouchers that can be issued (NULL for unlimited)',

    `current_issues_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Counter for the number of vouchers currently issued from this template',

    `max_issues_per_user` INT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'Maximum number of times a single user can claim a voucher from this template',

    `is_active` BOOLEAN NOT NULL DEFAULT TRUE COMMENT 'Indicates if this template is active for issuing new vouchers',

    `created_by_user_id` BIGINT UNSIGNED NULL COMMENT 'ID of the Admin/Staff user who created this template, FK to users(id)',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    INDEX `idx_vt_valid_until` (`valid_until`),
    INDEX `idx_vt_issuer` (`issuer_type`, `issuer_venue_id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Defines templates for various voucher campaigns, outlining their rules and conditions.';

-- table user_vouchers
CREATE TABLE `user_vouchers` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for an issued voucher',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT 'User ID who owns this voucher, FK to users(id)',
    `voucher_template_id` INT UNSIGNED NOT NULL COMMENT 'ID of the voucher template this instance is based on, FK to voucher_templates(id)',
    `voucher_code` VARCHAR(50) NOT NULL UNIQUE COMMENT 'Unique code for this voucher instance (can be auto-generated)',

    `issued_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp when this voucher was issued to the user',
    `expires_at` DATE NOT NULL COMMENT 'Expiry date for this specific voucher instance (copied from template or calculated)',

    `status` ENUM('AVAILABLE', 'USED', 'EXPIRED') NOT NULL DEFAULT 'AVAILABLE' COMMENT 'Status of this voucher',
    `used_at` DATETIME NULL COMMENT 'Timestamp when the voucher was used',
    `used_order_id` VARCHAR(255) NULL COMMENT 'Order ID where this voucher was used (optional)',
    `used_venue_id` INT UNSIGNED NULL COMMENT 'Venue ID where this voucher was used, FK to venues(id)',


    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    INDEX `idx_uv_user_status_expires` (`user_id`, `status`, `expires_at`),
    INDEX `idx_uv_code` (`voucher_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Specific voucher instances issued to users';

ALTER TABLE drunk_user 
MODIFY COLUMN create_time DATETIME(3),
MODIFY COLUMN update_time DATETIME(3),
MODIFY COLUMN last_login_time DATETIME(3);

