CREATE TABLE vehicles (
  id                      serial PRIMARY KEY,
  length                  numeric NOT NULL,
  passengers              integer NOT NULL,
  crew                    integer NOT NULL,
  name                    text NOT NULL,
  model                   text NOT NULL,
  manufacturer            text NOT NULL,
  cost_in_credits         text NOT NULL,
  max_atmosphering_speed  text NOT NULL,
  cargo_capacity          text NOT NULL,
  consumables             text NOT NULL,
  vehicle_class           text NOT NULL,
  url                     text NOT NULL,
  created                 timestamp NOT NULL DEFAULT CURRENT_DATE,
  edited                  timestamp NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE characters (
  id          serial PRIMARY KEY,
  height      integer NOT NULL,
  mass        integer NOT NULL,
  name        text NOT NULL,
  hair_color  text NOT NULL,
  skin_color  text NOT NULL,
  eye_color   text NOT NULL,
  birth_year  text NOT NULL,
  gender      text NOT NULL,
  homeworld   text NOT NULL,
  url         text NOT NULL,
  created     timestamp NOT NULL DEFAULT CURRENT_DATE,
  edited      timestamp NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE vehicles_pilots (
  pilot_id integer FOREIGN KEY REFERENCES characters(id),
  vehicle_id integer FOREIGN KEY REFERENCES vehicles(id),
  CONSTRAINT id PRIMARY KEY (pilot_id, vehicle_id)
)
