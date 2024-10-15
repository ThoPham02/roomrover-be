CREATE TABLE `contract_tbl` (
  `id` bigint,
  `code` varchar(255),
  `status` int,
  `renter_id` bigint,
  `renter_number` varchar(255),
  `renter_date` bigint,
  `renter_address` varchar(255),
  `renter_name` varchar(255),

  `lessor_id` bigint,
  `lessor_number` varchar(255),
  `lessor_date` bigint,
  `lessor_address` varchar(255),
  `lessor_name` varchar(255),

  `room_id` bigint,
  `check_in` bigint,
  `duration` int,
  `purpose` varchar(255),

  `created_at` bigint,
  `updated_at` bigint,
  `created_by` bigint,
  `updated_by` bigint,
  PRIMARY KEY (`id`)
);

CREATE TABLE `payment_tbl` (
  `id` bigint,
  `contract_id` bigint not null,
  `amount` bigint not null,
  `discount` bigint not null,
  `deposit` bigint not null,
  `deposit_date` bigint not null,
  `next_bill` bigint not null,
  PRIMARY KEY (`id`)
);

CREATE TABLE `payment_detail_tbl` (
  `id` bigint,
  `payment_id` bigint,
  `name` varchar(255),
  `type` int,
  `price` bigint,
  PRIMARY KEY (`id`)
);

CREATE TABLE `payment_renter_tbl` (
  `id` bigint,
  `payment_id` bigint,
  `user_id` bigint,
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

CREATE TABLE `bill_detail_tbl` (
  `id` bigint,
  `bill_id` bigint,
  `name` varchar(255),
  `price` bigint,
  `type` int,
  `quantity` int,
  `status` int,
  PRIMARY KEY (`id`)
);

CREATE TABLE `bill_pay_tbl` (
  `id` bigint,
  `bill_id` bigint,
  `amount` bigint,
  `pay_date` bigint,
  `url` varchar(255),
  `user_id` bigint,
  PRIMARY KEY (`id`)
);