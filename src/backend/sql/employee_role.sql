create role employee with
    login
    nosuperuser
    nocreatedb
    nocreaterole
    noreplication
    password '123'
    connection limit -1;

grant select, update, insert on
    employee,
    info_card,
    document,
    field,
    photo to employee;