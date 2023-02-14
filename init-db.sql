CREATE DATABASE AccountServiceApp;

USE AccountServiceApp;

CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext DEFAULT NULL,
  `phone_number` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP
  
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

CREATE TABLE `balances` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `total_balance` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_balances_user` (`user_id`),
  CONSTRAINT `fk_balances_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP
  
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

CREATE TABLE `balance_types` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `balance_type` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

CREATE TABLE `history_balances` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `user_id_to` bigint(20) unsigned DEFAULT NULL,
  `balance_type_id` bigint(20) unsigned DEFAULT NULL,
  `total` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_history_balances_user` (`user_id`),
  KEY `fk_history_balances_user_to` (`user_id_to`),
  KEY `fk_history_balances_balance_type` (`balance_type_id`),
  CONSTRAINT `fk_history_balances_balance_type` FOREIGN KEY (`balance_type_id`) REFERENCES `balance_types` (`id`),
  CONSTRAINT `fk_history_balances_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_history_balances_user_to` FOREIGN KEY (`user_id_to`) REFERENCES `users` (`id`),
  `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

INSERT INTO balance_types (balance_type)
VALUES ("Top Up"), ("Transfer");
