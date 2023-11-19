CREATE TABLE sensors
(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    group_id BIGINT REFERENCES sensor_groups(id),
    x NUMERIC,
    y NUMERIC,
    z NUMERIC,
    data_output_rate BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
