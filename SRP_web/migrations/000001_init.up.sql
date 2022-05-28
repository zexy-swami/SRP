create table if not exists users (
	user_id             smallserial,
	srp_id              varchar(100) not null,
	user_password_hash  varchar(100) not null,
	
	constraint user_id_pk    primary key (user_id),
	constraint unique_srp_id unique (srp_id)
);