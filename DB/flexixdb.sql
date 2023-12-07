-- Set character set and collation
SET character_set_client = 'utf8';
ALTER DATABASE flexixdb CHARACTER SET utf8 COLLATE utf8_general_ci;

-- Create the database
CREATE DATABASE IF NOT EXISTS flexixdb;
USE flexixdb;

-- Create the users table
CREATE TABLE IF NOT EXISTS `users` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Fname` varchar(255) DEFAULT NULL,
  `Lname` varchar(255) DEFAULT NULL,
  `Email` varchar(255) DEFAULT NULL,
  `Pass` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `Email` (`Email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- Insert data into the users table
INSERT INTO `users` (`ID`, `Fname`, `Lname`, `Email`, `Pass`) VALUES
(1, 'John', 'Dog', 'john.dog@example.com', 'password123'),
(2, 'Johnwick', 'Doggy', 'johnwick.doggy@example.com', 'password123');

-- Create the workspace table
CREATE TABLE IF NOT EXISTS `workspace` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Users_ID` int(11) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `Users_ID` (`Users_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- Insert data into the workspace table
INSERT INTO `workspace` (`ID`, `Users_ID`) VALUES
(1, 1),
(2, 2);

-- Create the project table
CREATE TABLE IF NOT EXISTS `project` (
  `ID` int(11) NOT NULL,
  `Project_name` varchar(255) DEFAULT NULL,
  `LDate_Time` datetime DEFAULT NULL,
  `Project_Path` varchar(255) DEFAULT NULL,
  `Screen_Img` varchar(255) DEFAULT NULL,
  `Workspace_ID` int(11) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `Workspace_ID` (`Workspace_ID`),
  CONSTRAINT `project_ibfk_1` FOREIGN KEY (`Workspace_ID`) REFERENCES `workspace` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- Set AUTO_INCREMENT for tables
ALTER TABLE `users` MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
ALTER TABLE `workspace` MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
