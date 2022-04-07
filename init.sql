CREATE TABLE IF NOT EXISTS orders (
    order_uid varchar PRIMARY KEY,
    track_number varchar,
    entry varchar
);

CREATE TABLE IF NOT EXISTS delivery (
    order_uid   varchar REFERENCES orders,
    name        varchar,
    phone       varchar,
    zip         varchar,
    city        varchar,
    address     varchar,
    region      varchar,
    email       varchar
);

CREATE TABLE IF NOT EXISTS payment (
    order_uid       varchar REFERENCES orders,
    transaction     varchar,
    request_id      varchar,
    currency        varchar,
    provider        varchar,
    amount          integer CHECK(amount > 0),
    payment_dt      integer CHECK(payment_dt > 0),
    bank            varchar,
    delivery_cost   integer CHECK(delivery_cost >= 0),
    good_total      integer CHECK(good_total > 0),
    custom_fee      integer CHECK(custom_fee >= 0)
);

CREATE TABLE IF NOT EXISTS items (
    order_uid       varchar REFERENCES orders,
    chrt_id         integer NOT NULL,
    track_number    varchar REFERENCES orders(track_number),
    price           integer CHECK(price > 0),
    rid             varchar,
    name            varchar,
    sale            integer CHECK(sale >= 0),
    size            varchar,
    total_price     integer CHECK(total_price > 0),
    nm_id           integer CHECK(nm_id >= 0),
    brand           varchar,
    status          integer
);
