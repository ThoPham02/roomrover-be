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