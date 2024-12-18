CREATE OR REPLACE FUNCTION create_role(_user TEXT, _pass TEXT)
RETURNS VOID AS $$
BEGIN
   IF EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = _user) THEN

      RAISE NOTICE 'Role "%" already exists. Skipping.', _user;
ELSE
    EXECUTE format('CREATE ROLE %I LOGIN PASSWORD %L', _user, _pass);
END IF;
END;
$$ LANGUAGE plpgsql;
--SEPARATOR
CREATE OR REPLACE FUNCTION create_database(_dbName TEXT, _owner TEXT)
RETURNS VOID AS $$
BEGIN
   CREATE EXTENSION IF NOT EXISTS dblink;
   IF EXISTS (
      SELECT FROM pg_database
      WHERE datname = _dbName) THEN

      RAISE NOTICE 'Database "%" already exists. Skipping.', _dbName;
ELSE
    PERFORM dblink_exec('dbname=' || current_database(), 'CREATE DATABASE ' || _dbName);
    EXECUTE format('ALTER DATABASE %I OWNER TO %I', _dbName, _owner);
END IF;
END;
$$ LANGUAGE plpgsql;
--SEPARATOR
CREATE OR REPLACE FUNCTION grant_permission(_user TEXT)
RETURNS VOID AS $$
BEGIN
   IF EXISTS (
      SELECT FROM pg_roles
      WHERE rolname = _user) THEN

      EXECUTE format('GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO %I', _user);
ELSE
    RAISE EXCEPTION 'Role "%" does not exist.', _user;
END IF;
END;
$$ LANGUAGE plpgsql;