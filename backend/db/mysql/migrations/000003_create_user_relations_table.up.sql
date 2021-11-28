CREATE TABLE IF NOT EXISTS `user_relations`(
  `id` bigint(20) PRIMARY KEY AUTO_INCREMENT,
  `follower_user_id` bigint(20) NOT NULL,
  `followee_user_id` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  unique key `unique_follower_followee` (`follower_user_id`, `followee_user_id`),
  key `idx_followee_follower` (`followee_user_id`, `follower_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
