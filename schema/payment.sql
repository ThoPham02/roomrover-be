CREATE TABLE `payment_tbl` (
    `id` BIGINT,
    `contract_id` BIGINT NOT NULL,
    `month` BIGINT NOT NULL,
	`total` BIGINT NOT NULL,
    `status` BIGINT NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
)

CREATE TABLE `payment_used_tbl` (
    `id` BIGINT,
    `payment_id` BIGINT NOT NULL,
    `service_id` BIGINT NOT NULL,
    `index` BIGINT NOT NULL,
    `status` BIGINT NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
)

CREATE TABLE `payment_detail_tbl` (
    `id` BIGINT,
    `payment_id` BIGINT NOT NULL,
    `amount` BIGINT NOT NULL,
    `type` BIGINT NOT NULL,
    `utl` VARCHAR(255) NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
)

