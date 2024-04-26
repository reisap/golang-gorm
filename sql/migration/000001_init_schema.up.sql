CREATE TABLE `campaign_images` (
                                   `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                                   `campaign_id` int(11) DEFAULT NULL,
                                   `file_name` varchar(255) DEFAULT NULL,
                                   `is_primary` tinyint(4) DEFAULT NULL,
                                   `created_at` datetime DEFAULT NULL,
                                   `updated_at` datetime DEFAULT NULL,
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `campaigns` (
                             `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                             `user_id` int(11) DEFAULT NULL,
                             `name` varchar(255) DEFAULT NULL,
                             `short_description` varchar(255) DEFAULT NULL,
                             `description` text,
                             `perks` text,
                             `backer_count` int(11) DEFAULT NULL,
                             `goal_amount` int(11) DEFAULT NULL,
                             `current_amount` int(11) DEFAULT NULL,
                             `slug` varchar(255) DEFAULT NULL,
                             `created_at` datetime DEFAULT NULL,
                             `updated_at` datetime DEFAULT NULL,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `transactions` (
                                `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                                `campaign_id` int(11) DEFAULT NULL,
                                `user_id` int(11) DEFAULT NULL,
                                `amount` int(11) DEFAULT NULL,
                                `status` varchar(255) DEFAULT NULL,
                                `code` varchar(255) DEFAULT NULL,
                                `created_at` datetime DEFAULT NULL,
                                `updated_at` datetime DEFAULT NULL,
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `users` (
                         `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                         `name` varchar(255) DEFAULT NULL,
                         `occupation` varchar(255) DEFAULT NULL,
                         `email` varchar(255) DEFAULT NULL,
                         `password_hash` varchar(255) DEFAULT NULL,
                         `avatar_file_name` varchar(255) DEFAULT NULL,
                         `role` varchar(255) DEFAULT NULL,
                         `token` varchar(255) DEFAULT NULL,
                         `created_at` datetime DEFAULT NULL,
                         `updated_at` datetime DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;