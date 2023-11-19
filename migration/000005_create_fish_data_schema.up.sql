CREATE TABLE fish_data
(
    id BIGSERIAL PRIMARY KEY,
    species_id BIGINT REFERENCES fish_species(id),
    sensor_data_id BIGINT REFERENCES sensors_data(id),
    count BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);