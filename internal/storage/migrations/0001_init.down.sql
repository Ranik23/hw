--------------------------------------------------------
-- Drop Constraints
--------------------------------------------------------

-- Drop foreign key constraints
ALTER TABLE "library"."subscriptions" DROP CONSTRAINT IF EXISTS "FK_subscriptions_books";
ALTER TABLE "library"."subscriptions" DROP CONSTRAINT IF EXISTS "FK_subscriptions_subscribers";
ALTER TABLE "library"."m2m_books_genres" DROP CONSTRAINT IF EXISTS "FK_m2m_books_genres_books";
ALTER TABLE "library"."m2m_books_genres" DROP CONSTRAINT IF EXISTS "FK_m2m_books_genres_genres";
ALTER TABLE "library"."m2m_books_authors" DROP CONSTRAINT IF EXISTS "FK_m2m_books_authors_books";
ALTER TABLE "library"."m2m_books_authors" DROP CONSTRAINT IF EXISTS "FK_m2m_books_authors_authors";

-- Drop primary key constraints
ALTER TABLE "library"."m2m_books_genres" DROP CONSTRAINT IF EXISTS "PK_m2m_books_genres";
ALTER TABLE "library"."m2m_books_authors" DROP CONSTRAINT IF EXISTS "PK_m2m_books_authors";

-- Drop check constraint
ALTER TABLE "library"."subscriptions" DROP CONSTRAINT IF EXISTS "check_enum";

--------------------------------------------------------
-- Drop Tables
--------------------------------------------------------

-- Drop tables in reverse order of creation
DROP TABLE IF EXISTS "library"."subscriptions" CASCADE;
DROP TABLE IF EXISTS "library"."subscribers" CASCADE;
DROP TABLE IF EXISTS "library"."m2m_books_genres" CASCADE;
DROP TABLE IF EXISTS "library"."m2m_books_authors" CASCADE;
DROP TABLE IF EXISTS "library"."genres" CASCADE;
DROP TABLE IF EXISTS "library"."books" CASCADE;
DROP TABLE IF EXISTS "library"."authors" CASCADE;

--------------------------------------------------------
-- Drop Schema
--------------------------------------------------------

-- Drop the schema if it exists
DROP SCHEMA IF EXISTS "library" CASCADE;