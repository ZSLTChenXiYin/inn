CREATE DATABASE IF NOT EXISTS user CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `user`.users (
    user_id CHAR(36) PRIMARY KEY,               -- 用户ID（UUID）
    user_name VARCHAR(64) NOT NULL,             -- 用户名
    user_password_hash VARCHAR(255) NOT NULL,   -- 密码哈希
    email VARCHAR(128) NOT NULL UNIQUE,         -- 邮箱（唯一）
    personalized_signature VARCHAR(255) DEFAULT NULL,         -- 个性描述（允许为空）
    picture VARCHAR(255) DEFAULT NULL,          -- 头像（允许为空）
    created_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6), -- 创建时间
    updated_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6), -- 更新时间
    deleted_at TIMESTAMP(6) NULL DEFAULT NULL,  -- 删除时间
    INDEX idx_user_name (user_name),            -- 用户名索引
    INDEX idx_created_at (created_at),          -- 创建时间索引
    INDEX idx_deleted_at (deleted_at),          -- 删除时间索引
    INDEX idx_created_deleted (created_at, deleted_at), -- 复合索引
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
