/* MIGRATION UP */
CREATE TABLE failed_jobs (
    id int unsigned primary key,
    payload text,
    queue varchar(255),
    exception text,
    created_at TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
)