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