FROM mysql:8.0.23

RUN chown -R mysql:mysql /var/lib/mysql

COPY ./database/*.sql /docker-entrypoint-initdb.d/