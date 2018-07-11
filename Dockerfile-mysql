FROM circleci/mysql:8
COPY ./example/script/* /docker-entrypoint-initdb.d/
RUN mv /docker-entrypoint-initdb.d/gen.script.mysql.user.sql /docker-entrypoint-initdb.d/gen.before.sql
