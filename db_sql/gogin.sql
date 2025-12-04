-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.30 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for gogin
CREATE DATABASE IF NOT EXISTS `gogin` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `gogin`;

-- Dumping structure for table gogin.adm_menu
CREATE TABLE IF NOT EXISTS `adm_menu` (
  `id_menu` varchar(15) NOT NULL,
  `id_parent` varchar(15) DEFAULT NULL,
  `menu_name` varchar(50) NOT NULL,
  `controller` varchar(50) DEFAULT NULL,
  `class_icon` varchar(30) DEFAULT NULL,
  `menu_sort` bigint DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id_menu`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table gogin.adm_menu: ~7 rows (approximately)
INSERT INTO `adm_menu` (`id_menu`, `id_parent`, `menu_name`, `controller`, `class_icon`, `menu_sort`, `is_active`) VALUES
	('1.1', '', 'Home', 'home', '', 1, 1),
	('1.1.1', '1.1', 'Dashboard', 'dashboard', '', 2, 1),
	('1.2', '', 'Administrator', 'administrator', '', 3, 1),
	('1.2.1', '1.2', 'Administrator', 'administrator', '', 4, 1),
	('1.2.1.1', '1.2.1', 'Roles', 'roles', '', 5, 1),
	('1.2.1.2', '1.2.1', 'Menu Roles', 'menu-roles', '', 6, 1),
	('1.2.1.3', '1.2.1', 'Users', 'users', '', 7, 1);

-- Dumping structure for table gogin.adm_menu_roles
CREATE TABLE IF NOT EXISTS `adm_menu_roles` (
  `id_menu_role` bigint NOT NULL AUTO_INCREMENT,
  `id_role` bigint NOT NULL,
  `id_menu` varchar(15) NOT NULL,
  `is_view` tinyint(1) DEFAULT '1',
  `is_insert` tinyint(1) DEFAULT '1',
  `is_edit` tinyint(1) DEFAULT '1',
  `is_delete` tinyint(1) DEFAULT '1',
  `is_print` tinyint(1) DEFAULT '1',
  `is_approve` tinyint(1) DEFAULT '1',
  `is_active` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id_menu_role`),
  KEY `fk_menu_roles_role` (`id_role`),
  KEY `fk_menu_roles_menu` (`id_menu`),
  CONSTRAINT `fk_menu_roles_menu` FOREIGN KEY (`id_menu`) REFERENCES `adm_menu` (`id_menu`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_menu_roles_role` FOREIGN KEY (`id_role`) REFERENCES `adm_roles` (`id_role`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table gogin.adm_menu_roles: ~7 rows (approximately)
INSERT INTO `adm_menu_roles` (`id_menu_role`, `id_role`, `id_menu`, `is_view`, `is_insert`, `is_edit`, `is_delete`, `is_print`, `is_approve`, `is_active`) VALUES
	(1, 1, '1.1', 1, 1, 1, 1, 1, 1, 1),
	(2, 1, '1.1.1', 1, 1, 1, 1, 1, 1, 1),
	(3, 1, '1.2', 1, 1, 1, 1, 1, 1, 1),
	(4, 1, '1.2.1', 1, 1, 1, 1, 1, 1, 1),
	(5, 1, '1.2.1.1', 1, 1, 1, 1, 1, 1, 1),
	(6, 1, '1.2.1.2', 1, 1, 1, 1, 1, 1, 1),
	(7, 1, '1.2.1.3', 1, 1, 1, 1, 1, 1, 1);

-- Dumping structure for table gogin.adm_roles
CREATE TABLE IF NOT EXISTS `adm_roles` (
  `id_role` bigint NOT NULL AUTO_INCREMENT,
  `role_name` varchar(100) NOT NULL,
  `description` varchar(100) DEFAULT NULL,
  `created_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_role`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table gogin.adm_roles: ~2 rows (approximately)
INSERT INTO `adm_roles` (`id_role`, `role_name`, `description`, `created_date`, `updated_date`) VALUES
	(1, 'Admin', 'Admin', '2025-11-19 08:53:41', '2025-11-19 08:53:41'),
	(2, 'User', 'User', '2025-11-19 08:53:50', '2025-11-19 08:53:50');

-- Dumping structure for table gogin.adm_users
CREATE TABLE IF NOT EXISTS `adm_users` (
  `id_user` bigint NOT NULL AUTO_INCREMENT,
  `id_role` bigint NOT NULL,
  `username` varchar(255) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `refresh_token` varchar(250) DEFAULT NULL,
  `refresh_token_expired` datetime DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT '1',
  `created_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_date` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_user`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  KEY `fk_users_roles` (`id_role`),
  CONSTRAINT `fk_users_roles` FOREIGN KEY (`id_role`) REFERENCES `adm_roles` (`id_role`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table gogin.adm_users: ~0 rows (approximately)
INSERT INTO `adm_users` (`id_user`, `id_role`, `username`, `email`, `password`, `refresh_token`, `refresh_token_expired`, `is_active`, `created_date`, `updated_date`) VALUES
	(1, 1, 'Admin', 'admin@gmail.com', '$2a$12$.4GUG5AzzSdcIg81j6/RKeLoHgcNlPSN/GqnntRU4NujQdq9AsaeO', '', '0000-00-00 00:00:00', 1, '2025-11-19 15:54:17', '2025-11-19 15:54:24'),
	(2, 2, 'user', 'user@gmail.com', '$2a$14$8DTYWAtK6rKEfZ3ECR38feJWjYTDkBOCySex4naCZxRjQUdQ405Nq', '', '0000-00-00 00:00:00', 1, '2025-11-19 16:40:03', '2025-11-19 16:40:27');

-- Dumping structure for table gogin.schema_migrations
CREATE TABLE IF NOT EXISTS `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table gogin.schema_migrations: ~1 rows (approximately)
INSERT INTO `schema_migrations` (`version`, `dirty`) VALUES
	(5, 0);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
