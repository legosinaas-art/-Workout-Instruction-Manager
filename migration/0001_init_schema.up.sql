create table users (
    id            bigserial primary key,
    email         varchar(255) not null unique,
    username      varchar(100) not null unique,
    created_at    timestamptz not null default now()
);

create table workout_instructions (
    id            bigserial primary key,
    name         varchar(100) not null,
    notes         varchar(255),
    created_at    timestamptz not null default now()
);

create table exercises(
    id bigserial primary key,
    name varchar(100) not null,
    notes varchar(255),
    created_at timestamptz not null default now()
);
-- Таблица связи. linkTable для many to many отношений
create table workout_instructions_exercises(
    workout_instruction_id bigint,
    exercise_id bigint,
    order_num int,
    details jsonb,
    PRIMARY KEY (workout_instruction_id, exercise_id, order_num)    --составной ключ
);






create table stats(
    id bigint,
    user_id bigint not null,
    weight int,
    height int,
    created_at timestamptz not null default now()
);
create index idx_stats_user on stats(user_id);

create table meals(
    id bigint,
    user_id bigint,
    date date not null,
    title varchar(100) not null,
    kcal int not null,
    created_at timestamptz not null default now()
);
create index idx_meals_user_date on meals(user_id, date);