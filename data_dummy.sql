-- MySQL dump 10.13  Distrib 9.7.0, for Linux (x86_64)
--
-- Host: localhost    Database: scorely
-- ------------------------------------------------------
-- Server version	9.7.0

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
SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN= 0;

--
-- GTID state at the beginning of the backup 
--

SET @@GLOBAL.GTID_PURGED=/*!80000 '+'*/ '045b856a-5708-11f1-9871-461d8fceb157:1-69';

--
-- Table structure for table `answer_questions`
--

DROP TABLE IF EXISTS `answer_questions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `answer_questions` (
  `id_answer_question` bigint unsigned NOT NULL AUTO_INCREMENT,
  `option` varchar(1) DEFAULT NULL,
  `is_correct` tinyint(1) DEFAULT '0',
  `student_id` bigint unsigned DEFAULT NULL,
  `exam_id` bigint unsigned DEFAULT NULL,
  `exam_question_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id_answer_question`),
  KEY `fk_exam_questions_answer_question` (`exam_question_id`),
  KEY `fk_exams_answer_question` (`exam_id`),
  KEY `fk_students_answer_question` (`student_id`),
  CONSTRAINT `fk_exam_questions_answer_question` FOREIGN KEY (`exam_question_id`) REFERENCES `exam_questions` (`id_exam_question`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_exams_answer_question` FOREIGN KEY (`exam_id`) REFERENCES `exams` (`id_exam`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_students_answer_question` FOREIGN KEY (`student_id`) REFERENCES `students` (`id_student`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `answer_questions`
--

LOCK TABLES `answer_questions` WRITE;
/*!40000 ALTER TABLE `answer_questions` DISABLE KEYS */;
/*!40000 ALTER TABLE `answer_questions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `classes`
--

DROP TABLE IF EXISTS `classes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `classes` (
  `id_class` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `level_id` bigint unsigned DEFAULT NULL,
  `major_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id_class`),
  KEY `fk_levels_class` (`level_id`),
  KEY `fk_majors_class` (`major_id`),
  CONSTRAINT `fk_levels_class` FOREIGN KEY (`level_id`) REFERENCES `levels` (`id_level`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_majors_class` FOREIGN KEY (`major_id`) REFERENCES `majors` (`id_major`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `classes`
--

LOCK TABLES `classes` WRITE;
/*!40000 ALTER TABLE `classes` DISABLE KEYS */;
INSERT INTO `classes` VALUES (1,'X RPL 1',1,1,NULL,NULL),(2,'XI TKJ 1',2,2,NULL,NULL);
/*!40000 ALTER TABLE `classes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exam_questions`
--

DROP TABLE IF EXISTS `exam_questions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exam_questions` (
  `id_exam_question` bigint unsigned NOT NULL AUTO_INCREMENT,
  `question` text,
  `exam_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id_exam_question`),
  KEY `fk_exams_exam_question` (`exam_id`),
  CONSTRAINT `fk_exams_exam_question` FOREIGN KEY (`exam_id`) REFERENCES `exams` (`id_exam`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exam_questions`
--

LOCK TABLES `exam_questions` WRITE;
/*!40000 ALTER TABLE `exam_questions` DISABLE KEYS */;
INSERT INTO `exam_questions` VALUES (1,'HTML merupakan singkatan dari?',1,NULL,NULL),(2,'Tag untuk membuat link adalah?',1,NULL,NULL),(3,'CSS digunakan untuk?',1,NULL,NULL);
/*!40000 ALTER TABLE `exam_questions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exams`
--

DROP TABLE IF EXISTS `exams`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exams` (
  `id_exam` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name_exams` varchar(200) DEFAULT NULL,
  `dates` date DEFAULT NULL,
  `start_lesson` varchar(20) DEFAULT NULL,
  `end_lesson` varchar(20) DEFAULT NULL,
  `teacher_subject_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id_exam`),
  KEY `fk_teacher_subjects_exam` (`teacher_subject_id`),
  CONSTRAINT `fk_teacher_subjects_exam` FOREIGN KEY (`teacher_subject_id`) REFERENCES `teacher_subjects` (`id_teacher_subject`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exams`
--

LOCK TABLES `exams` WRITE;
/*!40000 ALTER TABLE `exams` DISABLE KEYS */;
INSERT INTO `exams` VALUES (1,'UTS Pemrograman Web','2026-07-01','08:00','10:00',1,NULL,NULL);
/*!40000 ALTER TABLE `exams` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `levels`
--

DROP TABLE IF EXISTS `levels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `levels` (
  `id_level` bigint unsigned NOT NULL AUTO_INCREMENT,
  `level` varchar(100) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id_level`),
  UNIQUE KEY `uni_levels_level` (`level`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `levels`
--

LOCK TABLES `levels` WRITE;
/*!40000 ALTER TABLE `levels` DISABLE KEYS */;
INSERT INTO `levels` VALUES (1,'X',NULL,NULL),(2,'XI',NULL,NULL),(3,'XII',NULL,NULL);
/*!40000 ALTER TABLE `levels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `majors`
--

DROP TABLE IF EXISTS `majors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `majors` (
  `id_major` bigint unsigned NOT NULL AUTO_INCREMENT,
  `major` varchar(50) DEFAULT NULL,
  `major_abbreviation` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id_major`),
  UNIQUE KEY `uni_majors_major` (`major`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `majors`
--

LOCK TABLES `majors` WRITE;
/*!40000 ALTER TABLE `majors` DISABLE KEYS */;
INSERT INTO `majors` VALUES (1,'Rekayasa Perangkat Lunak','RPL',NULL,NULL),(2,'Teknik Komputer Jaringan','TKJ',NULL,NULL);
/*!40000 ALTER TABLE `majors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `option_questions`
--

DROP TABLE IF EXISTS `option_questions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `option_questions` (
  `id_option_question` bigint unsigned NOT NULL AUTO_INCREMENT,
  `option` varchar(1) DEFAULT NULL,
  `description_option` longtext,
  `is_correct` tinyint(1) DEFAULT '0',
  `exam_question_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id_option_question`),
  KEY `fk_exam_questions_option_question` (`exam_question_id`),
  CONSTRAINT `fk_exam_questions_option_question` FOREIGN KEY (`exam_question_id`) REFERENCES `exam_questions` (`id_exam_question`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `option_questions`
--

LOCK TABLES `option_questions` WRITE;
/*!40000 ALTER TABLE `option_questions` DISABLE KEYS */;
INSERT INTO `option_questions` VALUES (1,'A','Hyper Text Markup Language',1,1,NULL,NULL),(2,'B','Home Tool Markup Language',0,1,NULL,NULL),(3,'C','Hyperlink Text Markup Language',0,1,NULL,NULL),(4,'D','Hyper Transfer Markup Language',0,1,NULL,NULL),(5,'A','<img>',0,2,NULL,NULL),(6,'B','<a>',1,2,NULL,NULL),(7,'C','<div>',0,2,NULL,NULL),(8,'D','<p>',0,2,NULL,NULL),(9,'A','Mengatur tampilan website',1,3,NULL,NULL),(10,'B','Membuat database',0,3,NULL,NULL),(11,'C','Membuat API',0,3,NULL,NULL),(12,'D','Mengelola server',0,3,NULL,NULL);
/*!40000 ALTER TABLE `option_questions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id_role` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name_role` varchar(191) NOT NULL,
  `code_role` varchar(191) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id_role`),
  UNIQUE KEY `uni_roles_name_role` (`name_role`),
  UNIQUE KEY `uni_roles_code_role` (`code_role`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'Admin','ADMIN',NULL,NULL),(2,'Teacher','TEACHER',NULL,NULL),(3,'Student','STUDENT',NULL,NULL);
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `students`
--

DROP TABLE IF EXISTS `students`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `students` (
  `id_student` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `nisn` varchar(10) DEFAULT NULL,
  `gender` varchar(10) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `phone` bigint DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `class_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id_student`),
  UNIQUE KEY `uni_students_nisn` (`nisn`),
  UNIQUE KEY `uni_students_phone` (`phone`),
  KEY `fk_users_student` (`user_id`),
  KEY `fk_classes_student` (`class_id`),
  CONSTRAINT `fk_classes_student` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id_class`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_users_student` FOREIGN KEY (`user_id`) REFERENCES `users` (`id_user`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `students`
--

LOCK TABLES `students` WRITE;
/*!40000 ALTER TABLE `students` DISABLE KEYS */;
INSERT INTO `students` VALUES (1,'Muhammad Rizky','1234567890','Laki-Laki','Bone',81222222222,3,1,NULL,NULL);
/*!40000 ALTER TABLE `students` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `subjects`
--

DROP TABLE IF EXISTS `subjects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `subjects` (
  `id_subject` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name_subject` varchar(100) DEFAULT NULL,
  `semester` varchar(100) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id_subject`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subjects`
--

LOCK TABLES `subjects` WRITE;
/*!40000 ALTER TABLE `subjects` DISABLE KEYS */;
INSERT INTO `subjects` VALUES (1,'Pemrograman Web','Ganjil',NULL,NULL),(2,'Basis Data','Genap',NULL,NULL);
/*!40000 ALTER TABLE `subjects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `teacher_subjects`
--

DROP TABLE IF EXISTS `teacher_subjects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `teacher_subjects` (
  `id_teacher_subject` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_teachers` bigint unsigned DEFAULT NULL,
  `id_subjects` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id_teacher_subject`),
  KEY `idx_teacher_subjects_id_teachers` (`id_teachers`),
  KEY `idx_teacher_subjects_id_subjects` (`id_subjects`),
  CONSTRAINT `fk_teacher_subjects_subject` FOREIGN KEY (`id_subjects`) REFERENCES `subjects` (`id_subject`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_teacher_subjects_teacher` FOREIGN KEY (`id_teachers`) REFERENCES `teachers` (`id_teacher`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teacher_subjects`
--

LOCK TABLES `teacher_subjects` WRITE;
/*!40000 ALTER TABLE `teacher_subjects` DISABLE KEYS */;
INSERT INTO `teacher_subjects` VALUES (1,1,1),(2,1,2);
/*!40000 ALTER TABLE `teacher_subjects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `teachers`
--

DROP TABLE IF EXISTS `teachers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `teachers` (
  `id_teacher` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `nip` varchar(50) DEFAULT NULL,
  `gender` varchar(10) DEFAULT NULL,
  `address` varchar(50) DEFAULT NULL,
  `phone` bigint DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id_teacher`),
  UNIQUE KEY `uni_teachers_nip` (`nip`),
  UNIQUE KEY `uni_teachers_phone` (`phone`),
  KEY `fk_users_teacher` (`user_id`),
  CONSTRAINT `fk_users_teacher` FOREIGN KEY (`user_id`) REFERENCES `users` (`id_user`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teachers`
--

LOCK TABLES `teachers` WRITE;
/*!40000 ALTER TABLE `teachers` DISABLE KEYS */;
INSERT INTO `teachers` VALUES (1,'Andi Saputra','198812312024001001','Laki-Laki','Makassar',81211111111,2,NULL,NULL);
/*!40000 ALTER TABLE `teachers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id_user` bigint unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(50) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id_user`),
  UNIQUE KEY `uni_users_email` (`email`),
  KEY `fk_roles_users` (`role_id`),
  CONSTRAINT `fk_roles_users` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id_role`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'test@test.com','$2a$12$s5A8TAIAWsVwFWE.Ag/IleQTeGjWWrPcFDqBs23WRFuliq9kl.ydO',1,'2026-06-26 01:04:42.739',NULL),(2,'test2@test.com','$2a$12$8PdJiJJQbp3emND/Hxzj1ux9DEgKpkAxD/APAH41RXRyhhMxlEUN.',2,'2026-06-26 01:05:09.350',NULL),(3,'test3@test.com','$2a$12$1EGTTXpbGX.kR46hqOuKK.B1Miu/ZtBZ5eKyhDJ9IofcA1ZgW2eh.',3,'2026-06-26 01:05:17.302',NULL),(7,'coba@tes.com','$2a$12$eE/whJ/p/T6Y0hoevdCeI.dQIVjQ20bwYFcafE9NBDIeMu8kJYVya',3,'2026-06-26 21:19:06.030','2026-06-26 21:19:18.618');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
SET @@SESSION.SQL_LOG_BIN = @MYSQLDUMP_TEMP_LOG_BIN;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-06-26 13:30:38
