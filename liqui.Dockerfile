FROM liquibase/liquibase:latest

WORKDIR /app

ADD db/* .

RUN lpm add mysql --global