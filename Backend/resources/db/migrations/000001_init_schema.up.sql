create table documents (
	id BIGSERIAL primary key,
	title VARCHAR(50),
	format VARCHAR(50),
	author VARCHAR(50),
	owner VARCHAR(50),
	validity DATE,
	deleted_at timestamp WITH TIME ZONE,
	created_at timestamp WITH TIME ZONE,
	updated_at timestamp WITH TIME ZONE 
);

create table users (
	id BIGSERIAL primary key,
	first_name VARCHAR(50),
	last_name VARCHAR(50),
	age INT,
	email VARCHAR(50),
	city VARCHAR(50),
	phone VARCHAR(50),
	birth_date DATE,
	sex VARCHAR(50),
	country VARCHAR(50),
	doc_id BIGINT references documents(id),
	deleted_at timestamp WITH TIME ZONE,
	created_at timestamp WITH TIME ZONE,
	updated_at timestamp WITH TIME ZONE 
);