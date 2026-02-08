-- Payment Service Schema
-- Manages payments, payment methods, and transactions

CREATE TABLE IF NOT EXISTS payment_methods (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID NOT NULL,  -- references user/auth service (external ID)
    type        VARCHAR(30) NOT NULL CHECK (type IN ('card', 'bank', 'wallet', 'cod', 'other')),
    provider    VARCHAR(50),
    last_four   VARCHAR(4),
    expiry      VARCHAR(7),  -- MM/YYYY
    is_default  BOOLEAN NOT NULL DEFAULT false,
    metadata    JSONB,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS payments (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id        UUID NOT NULL,  -- references order service (external ID)
    user_id         UUID NOT NULL,
    amount          DECIMAL(12, 2) NOT NULL CHECK (amount > 0),
    currency        VARCHAR(3) NOT NULL DEFAULT 'USD',
    status          VARCHAR(30) NOT NULL DEFAULT 'pending' CHECK (status IN (
        'pending', 'authorized', 'captured', 'failed', 'refunded', 'partially_refunded', 'cancelled'
    )),
    payment_method_id UUID REFERENCES payment_methods(id),
    gateway         VARCHAR(50),
    gateway_ref     VARCHAR(255),
    gateway_response JSONB,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS payment_transactions (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    payment_id   UUID NOT NULL REFERENCES payments(id) ON DELETE CASCADE,
    type         VARCHAR(30) NOT NULL CHECK (type IN ('auth', 'capture', 'refund', 'void', 'adjustment')),
    amount       DECIMAL(12, 2) NOT NULL,
    status       VARCHAR(20) NOT NULL CHECK (status IN ('pending', 'success', 'failed')),
    gateway_txn_id VARCHAR(255),
    gateway_response JSONB,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS refunds (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    payment_id   UUID NOT NULL REFERENCES payments(id) ON DELETE CASCADE,
    amount       DECIMAL(12, 2) NOT NULL CHECK (amount > 0),
    reason       VARCHAR(255),
    status       VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'completed', 'failed')),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_payments_order ON payments(order_id);
CREATE INDEX idx_payments_user ON payments(user_id);
CREATE INDEX idx_payments_status ON payments(status);
CREATE INDEX idx_payment_methods_user ON payment_methods(user_id);
CREATE INDEX idx_payment_transactions_payment ON payment_transactions(payment_id);
CREATE INDEX idx_refunds_payment ON refunds(payment_id);
