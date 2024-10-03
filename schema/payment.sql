CREATE TABLE `bill_detail_tbl` (
  `id` bigint,
  `bill_id` bigint,
  `name` varchar(255),
  `price` bigint,
  `type` int,
  `quantity` int,
  PRIMARY KEY (`id`)
);

CREATE TABLE `bill_tbl` (
  `id` bigint,
  `payment_id` bigint,
  `payment_date` bigint,
  `amount` bigint,
  `discount` bigint,
  `status` int,
  PRIMARY KEY (`id`)
);

CREATE TABLE `bill_pay_tbl` (
  `id` bigint,
  `bill_id` bigint,
  `amount` bigint,
  `pay_date` bigint,
  `user_id` bigint,
  PRIMARY KEY (`id`)
);