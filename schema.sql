CREATE TABLE administrators (
    id serial,
    username varchar(200) unique,
    password varchar(200)
);
/*password value for this hash is 'admin'*/
INSERT INTO administrators (id, username, password) VALUES
(1, 'admin', '$2a$04$hadDO4xaxSdMb7kf34PvyuzYVNGFLXO4synZ94Kf8g5kgMQCHPfBC');

CREATE TABLE customers (
    id serial primary key,
    name varchar(200),
    phone varchar(200)
);


CREATE TABLE customer_addresses (
   id serial primary key,
   customer_id int,
   address varchar(200),
   zipcode varchar(5)
);
