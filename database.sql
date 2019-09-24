-- Adminer 4.7.3 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';

SET NAMES utf8mb4;

CREATE DATABASE `liver` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `liver`;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `qq` bigint(20) NOT NULL,
  `group` bigint(20) NOT NULL,
  `grade` double NOT NULL,
  `next` datetime NOT NULL,
  `ban` tinyint(1) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- 2019-09-24 03:29:31
