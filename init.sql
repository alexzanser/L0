CREATE TABLE IF NOT EXISTS orders (
    order_uid varchar PRIMARY KEY,
    track_number varchar,
    entry varchar
);

CREATE TABLE IF NOT EXISTS delivery (
    order_uid   REFERENCES orders,
    name        varchar,
    phone       varchar,
    zip         varchar,
    city        varchar,
    address     varchar,
    region      varchar,
    email       varchar
);

CREATE TABLE IF NOT EXISTS payment (
    order_uid       REFERENCES orders,
    transaction     varchar,
    request_id      varchar,
    currency        varchar,
    provider        varchar,
    amount          integer CHECK(VALUE > 0),
    payment_dt      integer CHECK(VALUE > 0),
    bank            varchar,
    delivery_cost   integer CHECK(VALUE >= 0),
    good_total      integer CHECK(VALUE > 0),
    custom_fee      integer CHECK(VALUE >= 0)
);

CREATE TABLE IF NOT EXISTS items (
    order_uid       REFERENCES orders,
    chrt_id         integer NOT NULL,
    track_number    REFERENCES orders(track_number),
    price           integer CHECK(VALUE > 0),
    rid             varchar,
    name            varchar,
    sale            integer CHECK(VALUE >= 0),
    size            varchar,
    total_price     integer CHECK(VALUE > 0),
    nm_id           integer CHECK(VALUE >= 0),
    brand           varchar,
    status          integer
);
