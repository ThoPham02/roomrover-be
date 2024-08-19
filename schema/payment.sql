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
