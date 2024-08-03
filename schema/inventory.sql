CREATE TABLE `provinces` (
    `id` INT,
    `name` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `districts` (
    `id` INT,
    `name` VARCHAR(255) NOT NULL,
    `province_id` INT,
    PRIMARY KEY (`id`)
);

CREATE TABLE `wards` (
    `id` INT,
    `name` VARCHAR(255) NOT NULL,
    `district_id` INT,
    PRIMARY KEY (`id`)
);

CREATE TABLE `houses` (
    `id` BIGINT,
    `name` VARCHAR(255) NOT NULL COLLATE utf8mb4_unicode_ci,
    `type` INT NOT NULL,

    `specific_address` VARCHAR(255) NOT NULL COLLATE utf8mb4_unicode_ci,
    `ward_id` INT,
    `district_id` INT,
    `province_id` INT,

    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `house_services` (
    `service_id` BIGINT,
    `house_id` BIGINT,
    `name` VARCHAR(255) NOT NULL COLLATE utf8mb4_unicode_ci,
    `price` DOUBLE NOT NULL,
    `type` INT NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`service_id`)
);

CREATE TABLE `classes` (
    `id` BIGINT,
    `lenssor_id` BIGINT,
    `house_id` BIGINT,
    `name` VARCHAR(255) NOT NULL COLLATE utf8mb4_unicode_ci,
    `description` TEXT NOT NULL COLLATE utf8mb4_unicode_ci,
    `acreage` INT NOT NULL,
    `price` DOUBLE NOT NULL,
    `status` INT NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `class_albums` (
    `id` BIGINT,
    `class_id` BIGINT,
    `url` VARCHAR(255) NOT NULL,
    `type` INT NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `rooms` (
    `id` BIGINT,
    `class_id` BIGINT,
    `name` VARCHAR(255) NOT NULL COLLATE utf8mb4_unicode_ci,
    `status` INT NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
);
