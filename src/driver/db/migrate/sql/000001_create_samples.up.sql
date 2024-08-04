create table samples (
  id char(26) not null primary key,
  name varchar(255) not null,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp,
  deleted_at datetime default null,
  index idx_deleted_at(deleted_at)
);
