CREATE DATABASE IF NOT EXISTS EnvoyABTesting;

CREATE TABLE IF NOT EXISTS EnvoyABTesting.Users 
(
  UserId INT NOT NULL PRIMARY KEY,
  UserName VARCHAR(50) NOT NULL,
  Password VARCHAR(50) NOT NULL,
  CONSTRAINT user_uniqueness UNIQUE (UserId, UserName)
);

INSERT INTO EnvoyABTesting.Users
(
 UserId, UserName, Password
)
VALUES
(
 1, 'Taro', MD5("nyanpass")
),
(
 2, 'Jiro', MD5("nyanpass")
),
(
 3, 'Saburo', MD5("nyanpass")
),
(
 4, 'UryuSakuno', MD5("nyanpass")
),
(
 5, 'SenaAiri', MD5("nyanpass")
),
(
 6, 'HatsushibaAi', MD5("nyanpass")
),
(
 7, 'TomosakaTsubame', MD5("nyanpass")
),
(
 8, 'ShinoharaHanako', MD5("nyanpass")
),
(
 9, 'InuiSana', MD5("nyanpass")
),
(
 10, 'AmahaMiu', MD5("nyanpass")
);