CREATE TABLE da_transactions (
  da_tx_id BIGINT GENERATED ALWAYS AS IDENTITY,
  da_calldata bytea NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE INDEX ON da_transactions(da_tx_id);
