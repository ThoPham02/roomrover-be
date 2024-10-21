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