create role security_employee with
    login
    nosuperuser
    nocreatedb
    nocreaterole
    noreplication
    password '123'
    connection limit -1;

grant select on
    all tables in schema public to security_employee;

grant insert, update, delete on
    info_card,
    passage to security_employee;