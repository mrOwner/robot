-- create "candles" table
CREATE TABLE `candles` (`uid` varchar(36) NOT NULL, `date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, `open` double NOT NULL, `close` double NOT NULL, `high` double NOT NULL, `low` double NOT NULL, `volume` int NULL) CHARSET utf8 COLLATE utf8_general_ci;
