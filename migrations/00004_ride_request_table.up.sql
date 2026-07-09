CREATE TABLE ride_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ride_id UUID NOT NULL REFERENCES rides(id) ON DELETE CASCADE,
    driver_id UUID NOT NULL REFERENCES users(id),
    status VARCHAR(10) NOT NULL DEFAULT 'sent'
        CHECK (status IN ('sent','accepted','rejected','expired')),
    sent_at TIMESTAMPTZ DEFAULT now(),
    responded_at TIMESTAMPTZ
);