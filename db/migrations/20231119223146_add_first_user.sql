-- +goose Up
-- +goose StatementBegin
select 'up SQL query';

insert into public.users
    (id, first_name, second_name, birthdate, password_hash)
values
    ('35df14eb-a4b8-4121-a85c-2bffabfde9d9', 'First', 'Second', now()::date, '$2a$14$niE25ndPuTQWDPfTekWDwuo97WRlAnFiBfnX.Ofl5EKuJHFsZdJ6.')
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
select 'down SQL query';

delete from public.users
where first_name = 'First'
;
-- +goose StatementEnd
