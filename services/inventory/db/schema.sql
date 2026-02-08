-- Inventory Service Schema
-- Manages stock, warehouses, and stock movements

CREATE TABLE IF NOT EXISTS warehouses (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(255) NOT NULL,
    code        VARCHAR(50) NOT NULL UNIQUE,
    address     TEXT,
    is_active   BOOLEAN NOT NULL DEFAULT true,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS stock (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id   UUID NOT NULL,  -- references product service (external ID)
    warehouse_id UUID NOT NULL REFERENCES warehouses(id) ON DELETE CASCADE,
    quantity     INT NOT NULL DEFAULT 0 CHECK (quantity >= 0),
    reserved     INT NOT NULL DEFAULT 0 CHECK (reserved >= 0),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(product_id, warehouse_id)
);

CREATE TABLE IF NOT EXISTS stock_movements (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stock_id     UUID NOT NULL REFERENCES stock(id) ON DELETE CASCADE,
    type         VARCHAR(20) NOT NULL CHECK (type IN ('in', 'out', 'adjust', 'transfer', 'reserve', 'release')),
    quantity     INT NOT NULL,
    reference_id UUID,  -- e.g. order_id, transfer_id
    reference_type VARCHAR(50),
    reason       VARCHAR(255),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS stock_reservations (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stock_id     UUID NOT NULL REFERENCES stock(id) ON DELETE CASCADE,
    order_id     UUID NOT NULL,  -- references order service
    quantity     INT NOT NULL CHECK (quantity > 0),
    expires_at   TIMESTAMPTZ NOT NULL,
    status       VARCHAR(20) NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'fulfilled', 'expired', 'cancelled')),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_stock_product ON stock(product_id);
CREATE INDEX idx_stock_warehouse ON stock(warehouse_id);
CREATE INDEX idx_stock_movements_stock ON stock_movements(stock_id);
CREATE INDEX idx_stock_movements_created ON stock_movements(created_at);
CREATE INDEX idx_stock_reservations_order ON stock_reservations(order_id);
CREATE INDEX idx_stock_reservations_status ON stock_reservations(status);
