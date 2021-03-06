create database user_center;

CREATE TABLE user_info(
    id INT(10) NOT NULL AUTO_INCREMENT,
    user_name VARCHAR(2048) NOT NULL COMMENT '用户名',
    pass_word VARCHAR(2048) NOT NULL COMMENT '密码',
    status TINYINT(1)  DEFAULT 1 COMMENT '1可用',
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (id))ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户信息';

-- DROP TABLE user_info ;
INSERT INTO user_info (user_name,pass_word,status) VALUES ("admin","123456",1);


create table books(
    id int not null auto_increment,
    title varchar(100) not null,
    author varchar(100) not null,
    price double(11,2) not null,
    sales int not null,
    stock int not null,
    img_path varchar(100),
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- select database();  查看当前使用哪个库
