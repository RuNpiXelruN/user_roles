--\set dbName {database name}
--\set dbUser {(service) user name}
--\set dbAdmin {admin user}
--\set dbPass {'password for (service) user'}

-- Local:
\set dbName postgres
\set dbUser postgres
\set dbAdmin postgres
\set dbPass 'password'


PREPARE get_dbName AS SELECT * FROM regexp_split_to_table(:'dbName', ',');
PREPARE get_dbUser AS SELECT * FROM regexp_split_to_table(:'dbUser', ',');
PREPARE get_dbAdmin AS SELECT * FROM regexp_split_to_table(:'dbAdmin', ',');
PREPARE get_dbPass AS SELECT * FROM regexp_split_to_table(:'dbPass', ',');
