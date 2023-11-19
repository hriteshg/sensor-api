CREATE TABLE sensors_data
(
    id BIGSERIAL PRIMARY KEY,
    temperature DOUBLE PRECISION NOT NULL,
    transparency BIGINT NOT NULL,
    data_output_rate BIGINT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);
