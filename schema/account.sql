CREATE TABLE `users` (
    `user_id` BIGINT,
    `phone` VARCHAR(255) UNIQUE NOT NULL,
    `password_hash` VARCHAR(255) NOT NULL,
    `role` BIGINT,
    `status` BIGINT NOT NULL DEFAULT 1, 
    `address` VARCHAR(255),
    `full_name` VARCHAR(255),
    `avatar_url` VARCHAR(255),
    `birthday` BIGINT,
    `gender` BIGINT,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    PRIMARY KEY (user_id)
);
