CREATE TABLE `user_tbl` (
  `id` bigint,
  `phone` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `role` int,
  `status` int NOT NULL,
  `address` varchar(255),
  `full_name` varchar(255),
  `avatar_url` varchar(255),
  `birthday` bigint,
  `gender` int,
  `CCCD_number` varchar(255),
  `CCCD_date` bigint,
  `CCCD_address` varchar(255),
  `created_at` bigint,
  `updated_at` bigint,
  PRIMARY KEY (`id`)
);

CREATE TABLE `house_tbl` (
  `id` bigint,
  `user_id` bigint NOT NULL,
  `name` varchar(255),
  `description` text,
  `type` int NOT NULL,
  `area` int NOT NULL,
  `price` int NOT NULL,
  `status` int NOT NULL,
  `bed_num` int,
  `living_num` int,
  `address` varchar(255),
  `ward_id` int NOT NULL,
  `district_id` int NOT NULL,
  `province_id` int NOT NULL,
  `created_at` bigint,
  `updated_at` bigint,
  `created_by` bigint,
  `updated_by` bigint,
  PRIMARY KEY (`id`)
);

CREATE TABLE `album_tbl` (
  `id` bigint,
  `house_id` bigint,
  `url` varchar(255),
  PRIMARY KEY (`id`)
);

CREATE TABLE `room_tbl` (
  `id` bigint,
  `house_id` bigint,
  `name` varchar(255),
  `status` int NOT NULL,
  `capacity` int,
  `e_index` int,
  `w_index` int,
  PRIMARY KEY (`id`)
);

CREATE TABLE `service_tbl` (
  `id` bigint,
  `house_id` bigint,
  `name` varchar(255),
  `price` bigint,
  `unit` int,
  PRIMARY KEY (`id`)
);

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

create table `notification_tbl` (
    `id` bigint,
    `sender` bigint not null,
    `receiver` bigint not null,
    `ref_id` bigint not null,
    `ref_type` int not null,
    `title` text not null,
    `description` text not null,
    `priority` int not null,
    `due_date` bigint not null,
    `status` int not null,
    `unread` int default 1 not null,
    `created_at` bigint not null,
    primary key (`id`)
);