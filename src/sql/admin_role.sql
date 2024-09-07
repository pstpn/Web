create role service_admin with
    login
    superuser
    createdb
    createrole
    replication
    password '123'
    connection limit -1;