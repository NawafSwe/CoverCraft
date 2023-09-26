DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'cover_craft') THEN
    CREATE DATABASE cover_craft;
END IF;
END $$;