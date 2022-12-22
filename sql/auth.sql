-- Active: 1671501400020@@34.64.186.156@3306@auth

CREATE TABLE user (
    `uid` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `activated` BOOLEAN DEFAULT(true),
    `birth` VARCHAR(255),
    `created_at` DATETIME DEFAULT NOW() NOT NULL,
    `created_by` VARCHAR(255),
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `email_confirm` BOOLEAN DEFAULT(true),
    `first_name` VARCHAR(255),
    `gender` INT,
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `image_url` VARCHAR(255),
    `lang_key` VARCHAR(255),
    `last_modified` DATETIME,
    `last_name` VARCHAR(255),
    `nick_name` VARCHAR(255),
    `nickname_changed` BOOLEAN DEFAULT(false),
    `nonce` DECIMAL(19,2),
    `password` VARCHAR(255),
    `role` VARCHAR(255)
);

CREATE TABLE wallet (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` BIGINT NOT NULL,
    Foreign Key(`user_id`) REFERENCES user(`uid`),
    `wit` INT DEFAULT(300),
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME DEFAULT NOW() NOT NULL
);

CREATE TABLE ItemOwned (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `owner_id` INT NOT NULL,
    Foreign Key (`owner_id`) REFERENCES wallet(`id`),
    `nft_id` INT NOT NULL,
    `amount` INT, --- NFT 보유량
    `created_at` DATETIME NOT NULL
);

CREATE TABLE Access (
    `user_id` BIGINT NOT NULL,
    Foreign Key (`user_id`) REFERENCES user(`uid`),
    `access_time` DATETIME DEFAULT NOW(),
    `block_id` INT NOT NULL
)