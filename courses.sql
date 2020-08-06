-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: db
-- Generation Time: May 07, 2020 at 01:39 PM
-- Server version: 8.0.19
-- PHP Version: 7.4.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `courses`
--

-- --------------------------------------------------------

--
-- Table structure for table `courses`
--

CREATE TABLE `courses` (
  `id` int NOT NULL,
  `name` varchar(100) NOT NULL DEFAULT '',
  `description` varchar(100) NOT NULL DEFAULT '',
  `price` decimal(10,0) NOT NULL,
  `duration` varchar(100) NOT NULL DEFAULT '',
  `program` varchar(100) NOT NULL DEFAULT '',
  `img` varchar(100) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `courses`
--

INSERT INTO `courses` (`id`, `name`, `description`, `price`, `duration`, `program`, `img`) VALUES
(1, 'Курс по Vue.js', 'Полный курс по vue.js', '150', '120', '<ul>\r\n  <li>Coffee</li>\r\n  <li>Tea</li>\r\n  <li>Milk</li>\r\n</ul>', 'vue.jpg'),
(2, 'Курс по Go', 'Полный курс по Go', '200', '100', '', 'go_lang.jpg'),
(3, 'High Loads', 'Полный курс High Loads', '180', '180', '', 'go_lang.jpg'),
(4, 'Курс по Html', 'Полный курс по Html', '90', '80', '', 'go_lang.jpg'),
(5, 'Курс по CSS', 'Полный курс по CSS', '200', '100', '', 'go_lang.jpg'),
(6, 'Курс по SEO', 'Полный курс по SEO', '180', '180', '', 'go_lang.jpg');

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` int NOT NULL,
  `name` varchar(100) NOT NULL DEFAULT '',
  `email` varchar(100) NOT NULL DEFAULT '',
  `tel` varchar(100) NOT NULL DEFAULT '',
  `date_create` varchar(100) NOT NULL DEFAULT '',
  `password` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `name`, `email`, `tel`, `date_create`, `password`) VALUES
(1, 'Liza', 'lotus@gmail.com', '78639863865', '', '');

-- --------------------------------------------------------

--
-- Table structure for table `customer_course`
--

CREATE TABLE `customer_course` (
  `id` int NOT NULL,
  `id_customer` int NOT NULL,
  `id_course` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `customer_course`
--

INSERT INTO `customer_course` (`id`, `id_customer`, `id_course`) VALUES
(3, 1, 1),
(4, 1, 2);

-- --------------------------------------------------------

--
-- Table structure for table `customer_course_video`
--

CREATE TABLE `customer_course_video` (
  `id` int NOT NULL,
  `id_customer` int NOT NULL,
  `id_course` int NOT NULL,
  `current_video` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `customer_course_video`
--

INSERT INTO `customer_course_video` (`id`, `id_customer`, `id_course`, `current_video`) VALUES
(1, 1, 1, 2),
(2, 1, 2, 5);

-- --------------------------------------------------------

--
-- Table structure for table `goose_db_version`
--

CREATE TABLE `goose_db_version` (
  `id` bigint UNSIGNED NOT NULL,
  `version_id` bigint NOT NULL,
  `is_applied` tinyint(1) NOT NULL,
  `tstamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `goose_db_version`
--

INSERT INTO `goose_db_version` (`id`, `version_id`, `is_applied`, `tstamp`) VALUES
(1, 0, 1, '2020-05-07 13:18:53'),
(2, 20200326115433, 1, '2020-05-07 13:18:53'),
(3, 20200421155737, 1, '2020-05-07 13:18:53'),
(4, 20200421160346, 1, '2020-05-07 13:18:53'),
(5, 20200424160122, 1, '2020-05-07 13:18:53'),
(6, 20200428155847, 1, '2020-05-07 13:18:53'),
(7, 20200428155907, 1, '2020-05-07 13:18:53'),
(8, 20200506101729, 1, '2020-05-07 13:18:54');

-- --------------------------------------------------------

--
-- Table structure for table `questions`
--

CREATE TABLE `questions` (
  `id` int NOT NULL,
  `id_video` int NOT NULL,
  `text` varchar(100) NOT NULL DEFAULT '',
  `sequence` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `questions`
--

INSERT INTO `questions` (`id`, `id_video`, `text`, `sequence`) VALUES
(1, 1, 'WHat is your name?', 1),
(2, 1, 'WHat is your surname?', 2),
(3, 2, 'WHat is your surname 2?', 2),
(4, 2, 'WHat is your name 2?', 1),
(5, 10, 'WHat is your name golang5?', 1),
(6, 10, 'WHat is your surname golang5?', 2);

-- --------------------------------------------------------

--
-- Table structure for table `responses`
--

CREATE TABLE `responses` (
  `id` int NOT NULL,
  `id_question` int NOT NULL,
  `text` varchar(100) NOT NULL DEFAULT '',
  `correct` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `responses`
--

INSERT INTO `responses` (`id`, `id_question`, `text`, `correct`) VALUES
(1, 1, 'Liza', 0),
(2, 1, 'Lotus', 0),
(3, 1, 'Right', 1),
(4, 1, 'Bla', 0),
(5, 2, 'Bla Bla', 0),
(6, 2, 'Bla Bla Bla', 0),
(7, 2, 'Right', 1),
(8, 2, 'Wrong', 0),
(9, 4, 'Wrong', 0),
(10, 4, 'Some', 0),
(11, 4, 'New', 0),
(12, 4, 'Right', 1),
(13, 3, 'Right', 1),
(14, 3, 'wrong', 0),
(15, 3, 'bla bla', 0),
(16, 3, 'hello', 0),
(17, 5, '1', 0),
(18, 5, '2', 0),
(19, 5, '3', 0),
(20, 5, 'Right', 1),
(21, 6, 'Right', 1),
(22, 6, '1', 0),
(23, 6, '2', 0),
(24, 6, '3', 0);

-- --------------------------------------------------------

--
-- Table structure for table `videos`
--

CREATE TABLE `videos` (
  `id` int NOT NULL,
  `id_course` int NOT NULL,
  `title` varchar(100) NOT NULL DEFAULT '',
  `name` varchar(100) NOT NULL DEFAULT '',
  `sequence` int NOT NULL,
  `poster` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `videos`
--

INSERT INTO `videos` (`id`, `id_course`, `title`, `name`, `sequence`, `poster`) VALUES
(1, 1, 'Первое видео vue.js курс', 'vuejs1.mp4', 1, 'vue.jpg'),
(2, 1, 'Второе видео vue.js курс', 'vuejs2.mp4', 2, 'vue.jpg'),
(3, 2, 'Первое видео golang курс', 'golang1.mp4', 1, 'vue.jpg'),
(4, 2, 'Второе видео golang курс', 'golang2.mp4', 2, 'vue.jpg'),
(5, 1, 'Третее видео vue.js курс', 'vuejs3.mp4', 3, 'vue.jpg'),
(6, 1, 'Четвертое видео vue.js курс', 'vuejs4.mp4', 4, 'vue.jpg'),
(7, 1, 'Пятое видео vue.js курс', 'vuejs5.mp4', 5, 'vue.jpg'),
(8, 2, 'Третее видео golang курс', 'golang3.mp4', 3, 'vue.jpg'),
(9, 2, 'Четвертое видео golang курс', 'golang4.mp4', 4, 'vue.jpg'),
(10, 2, 'Пятое видео golang курс', 'golang5.mp4', 5, 'vue.jpg');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `courses`
--
ALTER TABLE `courses`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `customer_course`
--
ALTER TABLE `customer_course`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_customer` (`id_customer`),
  ADD KEY `id_course` (`id_course`);

--
-- Indexes for table `customer_course_video`
--
ALTER TABLE `customer_course_video`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_customer` (`id_customer`),
  ADD KEY `id_course` (`id_course`);

--
-- Indexes for table `goose_db_version`
--
ALTER TABLE `goose_db_version`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`);

--
-- Indexes for table `questions`
--
ALTER TABLE `questions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_video` (`id_video`);

--
-- Indexes for table `responses`
--
ALTER TABLE `responses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_question` (`id_question`);

--
-- Indexes for table `videos`
--
ALTER TABLE `videos`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_course` (`id_course`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `courses`
--
ALTER TABLE `courses`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `customer_course`
--
ALTER TABLE `customer_course`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `customer_course_video`
--
ALTER TABLE `customer_course_video`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `goose_db_version`
--
ALTER TABLE `goose_db_version`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `questions`
--
ALTER TABLE `questions`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `responses`
--
ALTER TABLE `responses`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- AUTO_INCREMENT for table `videos`
--
ALTER TABLE `videos`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `customer_course`
--
ALTER TABLE `customer_course`
  ADD CONSTRAINT `customer_course_ibfk_1` FOREIGN KEY (`id_customer`) REFERENCES `customers` (`id`),
  ADD CONSTRAINT `customer_course_ibfk_2` FOREIGN KEY (`id_course`) REFERENCES `courses` (`id`);

--
-- Constraints for table `customer_course_video`
--
ALTER TABLE `customer_course_video`
  ADD CONSTRAINT `customer_course_video_ibfk_1` FOREIGN KEY (`id_customer`) REFERENCES `customers` (`id`),
  ADD CONSTRAINT `customer_course_video_ibfk_2` FOREIGN KEY (`id_course`) REFERENCES `courses` (`id`);

--
-- Constraints for table `questions`
--
ALTER TABLE `questions`
  ADD CONSTRAINT `questions_ibfk_1` FOREIGN KEY (`id_video`) REFERENCES `videos` (`id`);

--
-- Constraints for table `responses`
--
ALTER TABLE `responses`
  ADD CONSTRAINT `responses_ibfk_1` FOREIGN KEY (`id_question`) REFERENCES `questions` (`id`);

--
-- Constraints for table `videos`
--
ALTER TABLE `videos`
  ADD CONSTRAINT `videos_ibfk_1` FOREIGN KEY (`id_course`) REFERENCES `courses` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
