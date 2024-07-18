CREATE TABLE `users` (
                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                         `avatar` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                         `name` varchar(24) COLLATE utf8mb4_unicode_ci NOT NULL,
                         `phone` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
                         `password` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `status` int(10) DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci