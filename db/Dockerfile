FROM postgres:14-alpine

# Set environment variables for PostgreSQL
ENV POSTGRES_USER=your_user
ENV POSTGRES_PASSWORD=your_password
ENV POSTGRES_DB=your_database

# Expose PostgreSQL port (default: 5432)
EXPOSE 5432

# Copy the SQL migration script to the container
COPY data/sql/001_create_todos_table.sql /docker-entrypoint-initdb.d/

# Optionally, if you have other migration scripts or seeds, copy them here as well.
# COPY data/migrations/*.sql /docker-entrypoint-initdb.d/
# COPY data/seeds/*.sql /docker-entrypoint-initdb.d/



