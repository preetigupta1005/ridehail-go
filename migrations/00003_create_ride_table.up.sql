CREATE TABLE rides (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    passenger_id UUID NOT NULL REFERENCES users(id),
    driver_id UUID REFERENCES users(id), 

    pickup_lat DOUBLE PRECISION NOT NULL,
    pickup_lng DOUBLE PRECISION NOT NULL,
    pickup_address TEXT,

    drop_lat DOUBLE PRECISION NOT NULL,
    drop_lng DOUBLE PRECISION NOT NULL,
    drop_address TEXT,

    status VARCHAR(15) NOT NULL DEFAULT 'requested'
        CHECK (status IN ('requested','accepted','ongoing','completed','cancelled')),

    fare_amount NUMERIC(10,2),
    distance_km NUMERIC(6,2),

    requested_at TIMESTAMPTZ DEFAULT now(),
    accepted_at TIMESTAMPTZ,
    started_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    cancelled_at TIMESTAMPTZ,
    cancelled_by VARCHAR(10) CHECK (cancelled_by IN ('driver','passenger')),
    cancellation_reason TEXT
);
