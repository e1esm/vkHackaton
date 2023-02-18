CREATE SCHEMA IF NOT EXISTS nft_tickets_app;
CREATE TABLE IF NOT EXISTS
    `events` (
                          `Id` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
                          `title` varchar(45) NOT NULL,
                          `Description` varchar(255) DEFAULT NULL,
                          PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci