CREATE TABLE db_latihan_gin_prod.tbl_users
(
    id varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password  varchar(255) NOT NULL,
    role  varchar(10) NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    CONSTRAINT tbl_users_pkey PRIMARY KEY (id)
)