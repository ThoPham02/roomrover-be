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
);

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
);

CREATE TABLE `room_tbl` (
    `id` BIGINT,
    `house_id` BIGINT NOT NULL,
    `name` VARCHAR(255) NOT NULL COLLATE utf8mb4_unicode_ci,
    `capacity` BIGINT NOT NULL DEFAULT 0,
    `remain` BIGINT NOT NULL DEFAULT 0,
    `status` BIGINT NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
);

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
);
