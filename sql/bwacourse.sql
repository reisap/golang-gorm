CREATE TABLE `users` (
  `id` bigserial PRIMARY KEY,
  `name` varchar(100) NOT NULL,
  `occupation` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password_hash` varchar(100) NOT NULL,
  `role` varchar(100) NOT NULL,
  `token` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `campaign` (
  `id` bigserial PRIMARY KEY,
  `user_id` int NOT NULL,
  `name` varchar(100) NOT NULL,
  `short_description` varchar(100) NOT NULL,
  `description` text NOT NULL,
  `goal_amount` int NOT NULL,
  `current_amount` int NOT NULL,
  `backer_count` int NOT NULL DEFAULT 0,
  `perks` text NOT NULL,
  `slug` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `campaign_images` (
  `id` bigserial PRIMARY KEY,
  `campaign_id` int NOT NULL,
  `file_name` varchar(100) NOT NULL,
  `is_primary` bool DEFAULT false,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `transactions` (
  `id` bigserial PRIMARY KEY,
  `user_id` int NOT NULL,
  `campaign_id` int NOT NULL,
  `amount` int NOT NULL,
  `status` varchar(100) NOT NULL,
  `code` varchar(50),
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

ALTER TABLE `campaign` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `campaign_images` ADD FOREIGN KEY (`campaign_id`) REFERENCES `campaign` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`campaign_id`) REFERENCES `campaign` (`id`);
