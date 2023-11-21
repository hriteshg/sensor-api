INSERT INTO sensor_groups (name)
VALUES
    ('alpha'),
    ('beta');

INSERT INTO sensors (name, group_id, x_coordinate, y_coordinate, z_coordinate, data_output_rate)
VALUES
    ('alpha 1', 1, 10.0, 20.0, 30.0, 10),
    ('alpha 2', 1, 15.0, 40.0, 80.0, 20),
    ('beta 1', 2, 10.0, 20.0, 30.0, 30),
    ('beta 2', 2, 40.0, 40.0, 60.0, 40);

INSERT INTO sensors_data (temperature, transparency, sensor_id)
VALUES
    (40, 10, 1),
    (50, 20, 2),
    (60, 30, 3),
    (70, 40, 4);

INSERT INTO fish_data (species_name, sensor_data_id, count)
VALUES
    ('Atlantic Cod', 1, 5),
    ('Pacific Cod', 2, 12),
    ('Arabian Cod', 3, 30),
    ('Arctic Cod', 4, 45);