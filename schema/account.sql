CREATE TABLE `users` (
    `user_id` BIGINT,
    `profile_id` BIGINT,
    `username` VARCHAR(255) UNIQUE NOT NULL,
    `password_hash` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) UNIQUE NOT NULL,
    `role` BIGINT,
    PRIMARY KEY (user_id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'users table'

CREATE TABLE `profiles` (
    `profile_id` BIGINT,
    `fullname` VARCHAR(255),
    `dob` BIGINT,
    `avatar_url` VARCHAR(255),
    `address` VARCHAR(255),
    `phone` VARCHAR(255),
    `created_at` BIGINT,
    `updated_at` BIGINT,
    PRIMARY KEY (`profile_id`)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'profile table'

