FROM migrate/migrate

# Set workdir
WORKDIR /

# Copy migration files
COPY ./migrations ./sql/migrations

ENTRYPOINT ["migrate", "-source", "file:///sql/migrations"]