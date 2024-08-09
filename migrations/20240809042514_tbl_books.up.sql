CREATE TABLE db_latihan_gin_prod.tbl_books
(
    id varchar(255) NOT NULL,
    title varchar(255) NOT NULL,
    author varchar(255) NOT NULL,
    description text NOT NULL,
    price int NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    CONSTRAINT tbl_books_pkey PRIMARY KEY (id)
)