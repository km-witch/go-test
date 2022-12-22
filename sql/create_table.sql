-- Active: 1671431378608@@127.0.0.1@3306@Item
-- 어디에 추가하지?
CREATE TABLE Collection (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name_ko` VARCHAR(30) NOT NULL,
    `name_en` VARCHAR(30) NOT NULL,
    `publisher_ko` VARCHAR(30) NOT NULL,
    `publisher_en` VARCHAR(30) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `opensea_url` VARCHAR(30),
    `twitter_url` VARCHAR(30),
    `discord_url` VARCHAR(30),
    `ww_url` VARCHAR(30)
);

CREATE TABLE Product (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `collection_id` INT NOT NULL,
    Foreign Key (`collection_id`) REFERENCES Collection(`id`),
    `thumbnail_url` VARCHAR(30),
    `contract` VARCHAR(30),
    `name_ko` VARCHAR(30) NOT NULL,
    `name_en` VARCHAR(30) NOT NULL,
    `description` VARCHAR(30) NOT NULL,
    `amount` int,
    `properties` VARCHAR(30) NOT NULL,
    `snap` VARCHAR(30),
    `metadata` VARCHAR(256),
    `message_role` VARCHAR(30),
    `message_amount` INT,
    `created_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` TIMESTAMP
);

CREATE TABLE Nft (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `product_id` INT NOT NULL,
    Foreign Key (`product_id`) REFERENCES Product(`id`),
    `wallet_id` VARCHAR(30) NOT NULL,
    -- Foreign Key (`wallet_id`) REFERENCES Wallet(`id`), address(0)
    `token_id` INT NOT NULL,
    `name_ko` VARCHAR(30) NOT NULL,
    `name_en` VARCHAR(30) NOT NULL,
    `description` VARCHAR(30) NOT NULL,
    `properties` VARCHAR(30) NOT NULL,
    `created_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` TIMESTAMP
);

CREATE TABLE Nft_trans (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `nft_method` VARCHAR(30) NOT NULL,
    `from` VARCHAR(30) NOT NULL,
    `to` VARCHAR(30) NOT NULL,
    `nft_id` INT NOT NULL,
    Foreign Key (`nft_id`) REFERENCES Nft(`id`),
    `created_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Sales (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `product_id` INT NOT NULL,
    Foreign Key (`product_id`) REFERENCES Product(`id`),
    `status` VARCHAR(30) NOT NULL,
    `type` VARCHAR(30),
    `sale_count` VARCHAR(30) NOT NULL,
    `won_price` INT,
    `dollar_price` INT,
    `created_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Sales_log (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `sales_id` INT NOT NULL,
    Foreign Key (`sales_id`) REFERENCES Sales(`id`),
    `user_id` VARCHAR(30) NOT NULL,
    Foreign Key (`user_id`) REFERENCES user(`uid`),
    `amount` INT,
    `won_price` INT,
    `dollar_price` INT,
    `type` VARCHAR(30),
    `nft_id` INT,
    `created_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Block (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `thema` VARCHAR(30),
    `user_id` VARCHAR(30) NOT NULL,
    Foreign Key (`user_id`) REFERENCES user(`uid`),
    `name` VARCHAR(30) NOT NULL,
    `created_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` TIMESTAMP
);

CREATE TABLE Obj (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` VARCHAR(30) NOT NULL,
    Foreign Key (`user_id`) REFERENCES user(`uid`),
    `nft_id` VARCHAR(30) NOT NULL,
    `location_type` VARCHAR(30),
    `location` INT,
    `pos` VARCHAR(30) NOT NULL,
    `rot` VARCHAR(30) NOT NULL,
    `created_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` TIMESTAMP,
    `created_user` VARCHAR(30) NOT NULL,
    `updated_user` VARCHAR(30) NOT NULL
);

CREATE TABLE Obj_message (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `obj_id` INT NOT NULL,
    Foreign Key (`obj_id`) REFERENCES Obj(`id`),
    `user_id` VARCHAR(30) NOT NULL,
    Foreign Key (`user_id`) REFERENCES user(`uid`),
    `created_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `message` VARCHAR(256),
    `updated_time` TIMESTAMP NOT NULL,
    `created_user` VARCHAR(30) NOT NULL,
    `updated_user` VARCHAR(30) NOT NULL
);

CREATE TABLE Access (
    `user_id` INT,
    Foreign Key (`user_id`) REFERENCES withauthdb.user(`uid`),
    `access_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `block_id` INT,
    Foreign Key (`block_id`) REFERENCES Block(`id`)
);

CREATE TABLE Wallet (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT,
    Foreign Key (`user_id`) REFERENCES withauthdb.user(`uid`),
    `nft_id` VARCHAR(256),
    `wit` INT,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP
);