CREATE TABLE products
(
    prd_id          SERIAL PRIMARY KEY,
    prd_name        VARCHAR(100)   NOT NULL,
    prd_price       DECIMAL(10, 2) NOT NULL,
    prd_description TEXT,
    prd_stock       INT            NOT NULL,
    prd_active      BOOLEAN        NOT NULL DEFAULT TRUE,
    prd_created_at  TIMESTAMP               DEFAULT CURRENT_TIMESTAMP,
    prd_updated_at  TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
)