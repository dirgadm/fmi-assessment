## Fungsional Requirement
1. Registration(email, password, name, phone)user object
2. Login(email, password string)token string
3. Upload Photo(userId int, file string) file string
4. Absensi

## Schema Database

```
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

INSERT INTO `user` VALUES (1, 'dirga@gmail.com', '$2a$10$TGH1JvcjgszNEbtkzu.EteFsBB21dIJ00mqFoUVdiVlgkAgp3dvBq', 'Dirga Meligo', '85319076822', '2023-02-12 23:53:15', '2023-02-12 23:53:15');
```


