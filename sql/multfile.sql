CREATE DATABASE IF NOT EXISTS multfile CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `multfile`.files (
    file_id CHAR(32) PRIMARY KEY,             -- 文件唯一标识（UUID）
    file_name VARCHAR(255) NOT NULL,          -- 文件名
    file_path VARCHAR(255) NOT NULL,          -- 文件路径
    owner_id CHAR(32) NOT NULL,               -- 文件所有者（UUID）
    access_level TINYINT DEFAULT 0,           -- 访问权限（0:私有, 1:公开，2:密码）
    file_password_hash VARCHAR(255) DEFAULT NULL, -- 文件密码哈希
    check_sum CHAR(32) NOT NULL,              -- 文件校验值（MD5）
    created_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6), -- 创建时间
    updated_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6), -- 更新时间
    deleted_at TIMESTAMP(6) NULL DEFAULT NULL, -- 删除时间
    INDEX idx_file_name (file_name),          -- 文件名索引
    INDEX idx_owner_id (owner_id),            -- 所有者索引
    INDEX idx_created_at (created_at),        -- 创建时间索引
    INDEX idx_owner_created (owner_id, created_at), -- 复合索引
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
