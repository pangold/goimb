CREATE DATABASE IF NOT EXISTS imb_database

USE imb_database

CREATE TABLE imb_users(

);

CREATE TABLE friends(
    id INT NOT NULL AUTO_INCREMENT,
    user1_id INT NOT NULL,
    user2_id INT NOT NULL,
    created_at,
    updated_at,
    PRIMARY_KEY(id)
);

CREATE TABLE groups(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(32) NOT NULL,
    pic VARCHAR(100),
    announcement VARCHAR(100),
    created_at,
    updated_at,
    PRIMARY_KEY(id)
);

CREATE TABLE group_members(
    id INT NOT NULL AUTO_INCREMENT,
    group_id INT NOT NULL,
    user_id INT NOT NULL,
    role INT, -- 0(default) ordinary member, 1: master, 2: administrator,
    created_at,
    updated_at,
    PRIMARY_KEY(id)
);

-- FOR MangoDB / Redis
-- Expire at 7 days?
-- CREATE TABLE im_friend_request(
--     id INT,
--     user_id INT,
--     target_id INT,
--     result BOOLEAN, -- 0: REQUEST, 1: ACCEPT, 2: REJECT
--     postscript VARCHAR(100), -- Who you are?
--     time TIME,
-- )