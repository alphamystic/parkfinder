-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: May 29, 2024 at 07:38 AM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `park_finder`
--

-- --------------------------------------------------------

--
-- Table structure for table `contactus`
--

CREATE TABLE `contactus` (
  `cuid` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `subject` varchar(255) NOT NULL,
  `message` text NOT NULL,
  `handled` tinyint(1) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `newsletter`
--

CREATE TABLE `newsletter` (
  `newslt_id` int(11) NOT NULL,
  `email` varchar(255) NOT NULL,
  `subscribed` tinyint(1) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `parks`
--

CREATE TABLE `parks` (
  `id` int(11) NOT NULL,
  `park_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `location` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `parks`
--

INSERT INTO `parks` (`id`, `park_id`, `name`, `location`, `description`, `created_at`, `updated_at`) VALUES
(2, '65b0f4dd-a5b1-438b-a26d-000247694938', 'Dandenong Ranges National Park', 'Dandeong', '<p>Dandenong Ranges National Park is a rainforest getaway on Melbourne&#39;s doorstep. This is a place of tranquil forest walks, quaint hilltop towns and charismatic animals. Conquer the famous 1000 Steps, discover Olinda Falls and enjoy stunning views over Melbourne and the Yarra Valley. Here you&#39;ll find steep volcanic hills covered in stands of the world&#39;s tallest flowering tree, the Mountain Ash. Living among the greenery are wallabies, lyrebirds, wombats and the Powerful Owl. Please ensure you leave your dogs at home to avoid disturbing the local wildlife. Take an energetic walk up the 1000 Steps from Ferntree Gully Picnic Area and learn about its poignant association with Australia&#39;s Second World War veterans and the Kokoda Track Campaign.</p>\r\n', '2024-05-28 12:45:08', '2024-05-28 12:45:08'),
(3, '51d6fe85-d146-4e25-b0e3-c1684ae0838d', 'Flinders Chase National Park', 'New Zealand', '<p>Western end of Kangaroo Island features magnificent coastal landscapes coupled with vast wilderness areas and diverse wildlife. Flinders Chase National Park is the home of the iconic Admiral&#39;s Arch with its colony of New Zealand fur seals and the truly Remarkable Rocks. Open every day except Christmas day.</p>\r\n', '2024-05-28 16:29:34', '2024-05-28 16:29:34'),
(4, '84ac8413-111f-4eee-ab17-ffec010090d6', 'Port Campbell National Park', 'South Ocean', '<p>The wild Southern Ocean has carved the Port Campbell National Park coastline into formations that are famous the world over - and earned it the nickname of the Shipwreck Coast. Drive the Great Ocean Road and see London Bridge, The Grotto, Loch Ard Gorge and the unmissable Twelve Apostles. The best way to see Port Campbell National Park is to take the Great Ocean Road. This is one of the world&#39;s most celebrated scenic drives &ndash; and its undoubted highlight is the Twelve Apostles, which tower 45 metres above the Southern Ocean. On a coastline renowned for its spectacular coastal formations it&rsquo;s easy to overlook London Bridge and The Grotto, which are both as awesome in their own way as the more famous Twelve Apostles.</p>\r\n', '2024-05-28 16:49:10', '2024-05-28 16:49:10');

-- --------------------------------------------------------

--
-- Table structure for table `reviews`
--

CREATE TABLE `reviews` (
  `id` int(11) NOT NULL,
  `review_id` varchar(255) NOT NULL,
  `park_id` varchar(255) NOT NULL,
  `userid` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `comment` text NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `reviews`
--

INSERT INTO `reviews` (`id`, `review_id`, `park_id`, `userid`, `username`, `comment`, `created_at`, `updated_at`) VALUES
(1, 'b72a8976-248c-43ac-ad1d-adaac0987c2c', '65b0f4dd-a5b1-438b-a26d-000247694938', 'd488e793-953e-4a2f-8ddb-6f74f04b553a', 'admin', 'Testing Review if and how it works', '2024-05-28 15:27:24', '2024-05-28 15:27:24'),
(2, '354147f2-1c09-4f50-bba0-5c037fed09e4', '65b0f4dd-a5b1-438b-a26d-000247694938', 'd488e793-953e-4a2f-8ddb-6f74f04b553a', 'admin', 'Testing Review if and how it works', '2024-05-28 15:27:35', '2024-05-28 15:27:35'),
(3, 'd495a6d5-8d30-4e08-8218-e19a26a0c325', '51d6fe85-d146-4e25-b0e3-c1684ae0838d', 'd488e793-953e-4a2f-8ddb-6f74f04b553a', 'admin', 'Amazing place...', '2024-05-28 16:30:06', '2024-05-28 16:30:06'),
(4, '1ce6bea2-6588-4d2f-8cbc-4a6e9672a10b', '51d6fe85-d146-4e25-b0e3-c1684ae0838d', 'd488e793-953e-4a2f-8ddb-6f74f04b553a', 'admin', 'Highly reccommend this place', '2024-05-28 16:30:22', '2024-05-28 16:30:22');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `userid` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `name` varchar(25) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `userid`, `role`, `phone`, `name`, `email`, `password`, `created_at`, `updated_at`) VALUES
(12, 'd488e793-953e-4a2f-8ddb-6f74f04b553a', 'ADMIN', '1234567', 'admin', 'admin@pf.com', '$2a$10$2JnjKTUJc5HdliheN9Xr3.intUflGKTpaOyuqtbk21AGvqJX3Y4i.', '2024-05-27 21:26:59', '2024-05-27 21:26:59');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `newsletter`
--
ALTER TABLE `newsletter`
  ADD PRIMARY KEY (`newslt_id`);

--
-- Indexes for table `parks`
--
ALTER TABLE `parks`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `reviews`
--
ALTER TABLE `reviews`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `newsletter`
--
ALTER TABLE `newsletter`
  MODIFY `newslt_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `parks`
--
ALTER TABLE `parks`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `reviews`
--
ALTER TABLE `reviews`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
