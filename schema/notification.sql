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
)