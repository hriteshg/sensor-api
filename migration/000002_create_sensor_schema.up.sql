CREATE TABLE sensors
(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    group_id BIGINT REFERENCES sensor_groups(id),
    x_coordinate NUMERIC,
    y_coordinate NUMERIC,
    z_coordinate NUMERIC,
    data_output_rate BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
