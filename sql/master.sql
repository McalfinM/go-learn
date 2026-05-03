-- public.permissions definition

-- Drop table

-- DROP TABLE public.permissions;

CREATE TABLE public.permissions (
	id serial4 NOT NULL,
	permission_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
	"name" text NOT NULL,
	created_at timestamp NULL DEFAULT now(),
	CONSTRAINT permissions_name_key UNIQUE (name),
	CONSTRAINT permissions_pkey PRIMARY KEY (id)
);


-- public.roles definition

-- Drop table

-- DROP TABLE public.roles;

CREATE TABLE public.roles (
	id serial4 NOT NULL,
	role_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
	"name" text NOT NULL,
	created_at timestamp NULL DEFAULT now(),
	CONSTRAINT roles_name_key UNIQUE (name),
	CONSTRAINT roles_pkey PRIMARY KEY (id)
);


-- public.rooms definition

-- Drop table

-- DROP TABLE public.rooms;

CREATE TABLE public.rooms (
	room_id uuid NOT NULL DEFAULT gen_random_uuid(),
	"name" varchar(100) NULL,
	"location" text NULL,
	status varchar(20) NULL DEFAULT 'available'::character varying,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	price_transit numeric(12, 2) NULL,
	price_daily numeric(12, 2) NULL,
	price_monthly numeric(12, 2) NULL,
	deposit_amount numeric(12, 2) NULL DEFAULT 100000,
	is_transit_available bool NULL DEFAULT false,
	is_daily_available bool NULL DEFAULT true,
	is_monthly_available bool NULL DEFAULT true,
	is_active bool NULL DEFAULT false,
	id bigserial NOT NULL,
	CONSTRAINT rooms_pkey PRIMARY KEY (room_id)
);


-- public.booking_rates definition

-- Drop table

-- DROP TABLE public.booking_rates;

CREATE TABLE public.booking_rates (
	rate_id uuid NOT NULL DEFAULT gen_random_uuid(),
	room_id uuid NULL,
	booking_type varchar(20) NULL,
	price numeric(12, 2) NULL,
	CONSTRAINT booking_rates_pkey PRIMARY KEY (rate_id),
	CONSTRAINT booking_rates_room_id_fkey FOREIGN KEY (room_id) REFERENCES public.rooms(room_id)
);


-- public.bookings definition

-- Drop table

-- DROP TABLE public.bookings;

CREATE TABLE public.bookings (
	booking_id uuid NOT NULL DEFAULT gen_random_uuid(),
	user_id uuid NULL,
	room_id uuid NULL,
	booking_type varchar(20) NULL,
	start_time timestamp NULL,
	end_time timestamp NULL,
	duration_hours int4 NULL,
	base_price numeric(12, 2) NULL,
	deposit_amount numeric(12, 2) NULL DEFAULT 100000,
	penalty_amount numeric(12, 2) NULL DEFAULT 0,
	status varchar(20) NULL DEFAULT 'pending'::character varying,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	id bigserial NOT NULL,
	guest_name text NOT NULL,
	guest_ktp text NOT NULL,
	CONSTRAINT bookings_pkey PRIMARY KEY (booking_id),
	CONSTRAINT bookings_room_id_fkey FOREIGN KEY (room_id) REFERENCES public.rooms(room_id)
);


-- public.devices definition

-- Drop table

-- DROP TABLE public.devices;

CREATE TABLE public.devices (
	device_id uuid NOT NULL DEFAULT gen_random_uuid(),
	room_id uuid NULL,
	device_name varchar(100) NULL,
	api_key text NULL,
	last_seen timestamp NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT devices_api_key_key UNIQUE (api_key),
	CONSTRAINT devices_pkey PRIMARY KEY (device_id),
	CONSTRAINT devices_room_id_fkey FOREIGN KEY (room_id) REFERENCES public.rooms(room_id) ON DELETE CASCADE
);
CREATE INDEX idx_device_api_key ON public.devices USING btree (api_key);


-- public.payments definition

-- Drop table

-- DROP TABLE public.payments;

CREATE TABLE public.payments (
	payment_id uuid NOT NULL DEFAULT gen_random_uuid(),
	booking_id uuid NULL,
	payment_type varchar(20) NULL,
	amount numeric(12, 2) NULL,
	"method" varchar(50) NULL,
	status varchar(20) NULL DEFAULT 'pending'::character varying,
	paid_at timestamp NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	id bigserial NOT NULL,
	CONSTRAINT payments_pkey PRIMARY KEY (payment_id),
	CONSTRAINT payments_booking_id_fkey FOREIGN KEY (booking_id) REFERENCES public.bookings(booking_id) ON DELETE CASCADE
);


-- public.role_permissions definition

-- Drop table

-- DROP TABLE public.role_permissions;

CREATE TABLE public.role_permissions (
	id serial4 NOT NULL,
	role_id int4 NULL,
	permission_id int4 NULL,
	created_at timestamp NULL DEFAULT now(),
	CONSTRAINT role_permissions_pkey PRIMARY KEY (id),
	CONSTRAINT role_permissions_permission_id_fkey FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE,
	CONSTRAINT role_permissions_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE
);


-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	"name" varchar(255) NULL,
	email varchar(255) NOT NULL,
	phone varchar(20) NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	id bigserial NOT NULL,
	user_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
	role_id int4 NULL,
	"password" varchar(255) NOT NULL,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_user_uuid_unique UNIQUE (user_uuid),
	CONSTRAINT fk_users_role FOREIGN KEY (role_id) REFERENCES public.roles(id)
);
CREATE INDEX idx_users_user_uuid ON public.users USING btree (user_uuid);


-- public.access_logs definition

-- Drop table

-- DROP TABLE public.access_logs;

CREATE TABLE public.access_logs (
	log_id uuid NOT NULL DEFAULT gen_random_uuid(),
	device_id uuid NULL,
	booking_id uuid NULL,
	access_time timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	status varchar(20) NULL,
	CONSTRAINT access_logs_pkey PRIMARY KEY (log_id),
	CONSTRAINT access_logs_booking_id_fkey FOREIGN KEY (booking_id) REFERENCES public.bookings(booking_id),
	CONSTRAINT access_logs_device_id_fkey FOREIGN KEY (device_id) REFERENCES public.devices(device_id)
);


-- public.access_tokens definition

-- Drop table

-- DROP TABLE public.access_tokens;

CREATE TABLE public.access_tokens (
	token_id uuid NOT NULL DEFAULT gen_random_uuid(),
	booking_id uuid NULL,
	"token" text NULL,
	expired_at timestamp NULL,
	is_used bool NULL DEFAULT false,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT access_tokens_pkey PRIMARY KEY (token_id),
	CONSTRAINT access_tokens_token_key UNIQUE (token),
	CONSTRAINT access_tokens_booking_id_fkey FOREIGN KEY (booking_id) REFERENCES public.bookings(booking_id) ON DELETE CASCADE
);


-- public.profiles definition

-- Drop table

-- DROP TABLE public.profiles;

CREATE TABLE public.profiles (
	id bigserial NOT NULL,
	profile_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
	user_id int8 NULL,
	full_name text NULL,
	ktp_number text NULL,
	phone text NULL,
	address text NULL,
	date_of_birth date NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL DEFAULT now(),
	CONSTRAINT profiles_pkey PRIMARY KEY (id),
	CONSTRAINT profiles_user_id_key UNIQUE (user_id),
	CONSTRAINT profiles_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);