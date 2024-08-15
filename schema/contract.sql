CREATE TABLE `contract_tbl` (
    `id` BIGINT,
    `renter_id` BIGINT NOT NULL,
    `lessor_id` BIGINT NOT NULL,
    `room_id` BIGINT NOT NULL,
    `status` BIGINT NOT NULL,

    `contract_url` VARCHAR(255),
    `description` TEXT NOT NULL COLLATE utf8mb4_unicode_ci,
    `start_date` BIGINT NOT NULL,
    `end_date` BIGINT NOT NULL,
    `type` BIGINT NOT NULL COMMENT '0: k coc, 1: coc',
    `deposit` BIGINT NOT NULL,
    `deadline` BIGINT NOT NULL,

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
    PRIMARY KEY (`id`)
)