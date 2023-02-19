CREATE SCHEMA IF NOT EXISTS nft_tickets_app;



CREATE TABLE IF NOT EXISTS `events` (
                          `Id` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
                          `Title` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL,
                          `Description` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci DEFAULT NULL,
                          `Date` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci DEFAULT NULL,
                          `Platform` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL,
                          `Image` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci DEFAULT NULL,
                          `Link` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci DEFAULT NULL,
                          `Org_id` int DEFAULT NULL,
                          PRIMARY KEY (`Id`),
                          KEY `FK_Org_idx` (`Org_id`),
                          CONSTRAINT `FK_Org` FOREIGN KEY (`Org_id`) REFERENCES `organizers` (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_unicode_ci;

CREATE TABLE IF NOT EXISTS `organizers` (
                              `Id` int NOT NULL,
                              `Name` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci DEFAULT NULL,
                              `Address` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci DEFAULT NULL,
                              PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_unicode_ci;


CREATE TABLE IF NOT EXISTS `tokens` (
                          `Id` int NOT NULL,
                          `Category` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci DEFAULT NULL,
                          `Event_id` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
                          PRIMARY KEY (`Id`,`Event_id`),
                          KEY `FK_Event_idx` (`Event_id`),
                          CONSTRAINT `FK_Event` FOREIGN KEY (`Event_id`) REFERENCES `events` (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_unicode_ci;


CREATE TABLE IF NOT EXISTS `white_list` (
                              `event_id` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
                              `user_id` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL,
                              `Role` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL,
                              PRIMARY KEY (`event_id`,`user_id`,`Role`),
                              UNIQUE KEY `Uniq` (`event_id`,`user_id`,`Role`),
                              KEY `FK_users_idx` (`user_id`),
                              CONSTRAINT `FK_events` FOREIGN KEY (`event_id`) REFERENCES `events` (`Id`),
                              CONSTRAINT `FK_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_unicode_ci;

CREATE TABLE IF NOT EXISTS `users` (
                         `Id` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL,
                         `Public_Hash` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL,
                         PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_unicode_ci;

