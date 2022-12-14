DROP TABLE IF EXISTS `listPassword`;
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `password` varchar(50) NOT NULL,
                         `identifyNumber` varchar(20) NOT NULL,
                         `fullName` varchar(200) NOT NULL,
                         `userName` varchar(50) not null,
                         `email` varchar(50) not null,
                         `phone` varchar(20) DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `deleted_at` timestamp NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `identifyNumber` (`identifyNumber`)

);

CREATE TABLE `listPassword` (
                                `id` int not null AUTO_INCREMENT PRIMARY key,
                                `password` varchar(200) not null,
                                `userName` varchar(50) not null

);

