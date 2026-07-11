ALTER TABLE driver_details DROP COLUMN current_lat;
ALTER TABLE driver_details DROP COLUMN current_lng;
ALTER TABLE driver_details ADD COLUMN current_location GEOGRAPHY(POINT, 4326);
CREATE INDEX idx_driver_location ON driver_details USING GIST (current_location);