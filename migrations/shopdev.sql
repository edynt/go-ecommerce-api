-- MySQL dump 10.13  Distrib 8.3.0, for Linux (aarch64)
--
-- Host: localhost    Database: shopdev
-- ------------------------------------------------------
-- Server version	8.3.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `shopdev`
--

/*!40000 DROP DATABASE IF EXISTS `shopdev`*/;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `shopdev` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `shopdev`;

--
-- Table structure for table `go_crm_user_v2`
--

DROP TABLE IF EXISTS `go_crm_user_v2`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `go_crm_user_v2` (
  `usr_id` int NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
  `usr_email` longtext NOT NULL COMMENT 'Email',
  `us_phone` longtext NOT NULL COMMENT 'Phone Number',
  `usr_username` longtext NOT NULL COMMENT 'Username',
  `usr_password` longtext NOT NULL COMMENT 'Password',
  `usr_created_at` int NOT NULL COMMENT 'Creation Time',
  `usr_updated_at` int NOT NULL COMMENT 'Update Time',
  `usr_create_ip_at` longtext NOT NULL COMMENT 'Creation IP',
  `usr_last_login_at` int NOT NULL COMMENT 'Last Login Time',
  `usr_last_login_ip_at` longtext NOT NULL COMMENT 'Last Login IP',
  `usr_login_times` int NOT NULL COMMENT 'Login Times',
  `usr_status` tinyint(1) NOT NULL COMMENT 'Status 1:enable, 0:disable, -1:deleted',
  PRIMARY KEY (`usr_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
-- Dump completed on 2025-06-01  9:06:30
