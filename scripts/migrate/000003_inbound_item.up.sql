CREATE TABLE inbound_item (
    id TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP(9) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(9) NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,

    CONSTRAINT inbound_item_pkey PRIMARY KEY ("id")
);
