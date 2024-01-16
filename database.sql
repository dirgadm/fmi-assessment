CREATE TABLE IF NOT EXISTS `user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NULL DEFAULT NULL,
  `password` varchar(255) NULL DEFAULT NULL,
  `name` varchar(255) NULL DEFAULT NULL,
  `phone` varchar(255) NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB;

INSERT INTO `user` VALUES (1, 'dirga@gmail.com', '$2a$14$7G8yOp3882QCUCQiVg1tVOIbiZ.9bEEjPuAFxBNozqaq8FqSjHlqS', 'Dirga Meligo', '85319076822', '2023-02-12 23:53:15', '2023-02-12 23:53:15');

CREATE TABLE IF NOT EXISTS `uploaded_files` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `file_name` VARCHAR(255) NOT NULL,
    `user_id` INT,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
);

CREATE TABLE IF NOT EXISTS `attendance` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `latitude` decimal(10, 8) NOT NULL,
  `longitude` decimal(11, 8) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB;

DELIMITER //

CREATE PROCEDURE `absensi`(
  IN p_user_id INT,
  IN p_latitude DECIMAL(10, 8),
  IN p_longitude DECIMAL(11, 8)
)
BEGIN
  INSERT INTO `attendance` (`user_id`, `latitude`, `longitude`, `created_at`)
  VALUES (p_user_id, p_latitude, p_longitude, CURRENT_TIMESTAMP);
END //

DELIMITER ;

DELIMITER //

CREATE PROCEDURE `register_user`(
    IN p_email VARCHAR(255),
    IN p_password VARCHAR(255),
    IN p_name VARCHAR(255),
    IN p_phone VARCHAR(255)
)
BEGIN
    INSERT INTO `user` (`email`, `password`, `name`, `phone`)
    VALUES (p_email, p_password, p_name, p_phone);
END //

DELIMITER ;

DELIMITER //

CREATE PROCEDURE `register_user_with_photo`(
    IN p_email VARCHAR(255),
    IN p_password VARCHAR(255),
    IN p_name VARCHAR(255),
    IN p_phone VARCHAR(255),
    IN p_photo_filename VARCHAR(255)
)
BEGIN
    DECLARE last_user_id INT;

    -- Insert into user table
    INSERT INTO `user` (`email`, `password`, `name`, `phone`)
    VALUES (p_email, p_password, p_name, p_phone);

    -- Get the last inserted user id
    SELECT LAST_INSERT_ID() INTO last_user_id;

    -- Insert into uploaded_files table
    INSERT INTO `uploaded_files` (`file_name`, `user_id`)
    VALUES (p_photo_filename, last_user_id);
END //

DELIMITER ;

