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