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

CREATE TABLE `house_tbl` (
    `id` BIGINT,
    `user_id` BIGINT NOT NULL,
    `name` VARCHAR(255) NOT NULL COLLATE utf8mb4_unicode_ci,
    `description` TEXT NOT NULL COLLATE utf8mb4_unicode_ci,
    `util` INT NOT NULL,
    `type` BIGINT NOT NULL,

    `area` BIGINT NOT NULL,
    `price` BIGINT NOT NULL,
    `status` BIGINT NOT NULL,

    `address` VARCHAR(255) NOT NULL COLLATE utf8mb4_unicode_ci,
    `ward_id` BIGINT NOT NULL,
    `district_id` BIGINT NOT NULL,
    `province_id` BIGINT NOT NULL,

    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
)

CREATE TABLE `album_tbl` (
    `id` BIGINT,
    `house_id` BIGINT NOT NULL,
    `url` VARCHAR(255) NOT NULL,
    `type` BIGINT NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
)

CREATE TABLE `room_tbl` (
    `id` BIGINT,
    `house_id` BIGINT NOT NULL,
    `name` VARCHAR(255) NOT NULL COLLATE utf8mb4_unicode_ci,
    `status` BIGINT NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
)

CREATE TABLE `service_tbl` (
    `id` BIGINT,
    `house_id` BIGINT NOT NULL,
    `name` VARCHAR(255) NOT NULL COLLATE utf8mb4_unicode_ci,
    `price` BIGINT NOT NULL,
    `type` BIGINT NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
)

CREATE TABLE `contract_tbl` (
    `id` BIGINT,
    `room_id` BIGINT NOT NULL,
    `status` BIGINT NOT NULL,

    `contract_url` VARCHAR(255),
    `description` TEXT NOT NULL COLLATE utf8mb4_unicode_ci,
    `start` BIGINT NOT NULL,
    `end` BIGINT NOT NULL,
    `next_bill` BIGINT NOT NULL,

    `type` BIGINT NOT NULL COMMENT '0: k coc, 1: coc',
    `deposit` BIGINT NOT NULL,
    `deadline` BIGINT NOT NULL,
    `deposit_url` VARCHAR(255),

    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
)

CREATE TABLE `contract_renter_tbl` (
    `id` BIGINT,
    `contract_id` BIGINT NOT NULL,
    `renter_id` BIGINT NOT NULL,
    `type` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
)

CREATE TABLE `contract_detail_tbl` (
    `id` BIGINT,
    `contract_id` BIGINT NOT NULL,
    `service_id` BIGINT NOT NULL,
    `price` BIGINT NOT NULL,
    `type` BIGINT NOT NULL,
    `index` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
)

CREATE TABLE `bill` (
    `id` BIGINT,
    `contract_id` BIGINT NOT NULL,
    `total` BIGINT NOT NULL,
    `paid` BIGINT NOT NULL,
    `status` BIGINT NOT NULL,
    `month` BIGINT NOT NULL,

    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
)

CREATE TABLE `bill_detail` (
    `id` BIGINT,
    `bill_id` BIGINT NOT NULL,
    `contract_service_id` BIGINT NOT NULL,
    `price` BIGINT NOT NULL,
    `type` BIGINT NOT NULL,
    `quantity` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
)

CREATE TABLE `payment` (
    `id` BIGINT,
    `bill_id` BIGINT NOT NULL,
    `amount` BIGINT NOT NULL,
    `type` BIGINT NOT NULL,
    `url` VARCHAR(255) NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
)
