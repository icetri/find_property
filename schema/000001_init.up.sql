CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table adverts
(
    id                integer                                    not null
        constraint advert_pk
            primary key,
    subcategory       varchar(255)                               not null,
    rooms_count       varchar(255)                               not null,
    before_metro      boolean      default false                 not null,
    url               varchar(255)                               not null,
    price             bigint                                     not null,
    time_publish      text                                       not null,
    phone             varchar(255)                               not null,
    address           text         default ''::character varying not null,
    metro             text         default ''::character varying not null,
    offer_type        varchar(255) default ''::character varying not null,
    floor_number      integer      default 0                     not null,
    area              varchar(255) default ''::character varying not null,
    building_year     integer      default 0                     not null,
    is_apartments     boolean      default false                 not null,
    deal_type         varchar(255) default ''::character varying not null,
    repair_type       varchar(255) default ''::character varying not null,
    images            text         default ''::character varying not null,
    balconies         boolean      default false                 not null,
    loggias           boolean      default false                 not null,
    jk_fullurl        text         default ''::character varying not null,
    jk_name           text         default ''::character varying not null,
    published_user_id bigint       default 0                     not null,
    cian_user_id      bigint       default 0                     not null,
    floors_count      integer      default 0                     not null,
    decoration        varchar(255) default ''::character varying not null,
    developer_name    varchar(255) default ''::character varying not null,
    web_site_url_utm  text         default ''::character varying not null,
    lat               varchar(255) default ''::character varying not null,
    lng               varchar(255) default ''::character varying not null,
    deadline_info     text         default ''::character varying not null,
    jk_id             integer      default 0                     not null
);

create unique index advert_id_uindex
    on adverts (id);


create table filters
(
    id             uuid         default uuid_generate_v4()                                                                                                                                                                                                                                                                                                                                                                                                                not null
        constraint filters_pk
            primary key,
    subcategory    varchar(255) default ''::character varying                                                                                                                                                                                                                                                                                                                                                                                                  not null,
    rooms_count    text[]       default ARRAY ['1'::text]                                                                                                                                                                                                                                                                                                                                                                                                  not null,
    price1         bigint       default 7500000                                                                                                                                                                                                                                                                                                                                                                                                                           not null,
    price2         bigint       default 9000000                                                                                                                                                                                                                                                                                                                                                                                                                           not null,
    area1          varchar(255) default '40'::character varying                                                                                                                                                                                                                                                                                                                                                                                                           not null,
    area2          varchar(255) default '60'::character varying                                                                                                                                                                                                                                                                                                                                                                                                           not null,
    before_metro   varchar(255) default ''::character varying                                                                                                                                                                                                                                                                                                                                                                                                      not null,
    building_year1 integer      default 2010                                                                                                                                                                                                                                                                                                                                                                                                                              not null,
    building_year2 integer      default 2025                                                                                                                                                                                                                                                                                                                                                                                                                              not null,
    jk_name        varchar(255) default ''::character varying                                                                                                                                                                                                                                                                                                                                                                                                             not null,
    category       varchar(255) default ''::character varying                                                                                                                                                                                                                                                                                                                                                                                       not null,
    type_balcony   varchar(255) default ''::character varying                                                                                                                                                                                                                                                                                                                                                                                                      not null,
    repair_type    varchar(255) default ''::character varying                                                                                                                                                                                                                                                                                                                                                                                                      not null,
    floors         varchar(255) default ''::character varying                                                                                                                                                                                                                                                                                                                                                                                             not null,
    user_id        bigint                                                                                                                                                                                                                                                                                                                                                                                                                                                 not null,
    address        varchar(255) default ''::character varying                                                                                                                                                                                                                                                                                                                                                                                                             not null,
    "default"      text         default ''::character varying                                                                                                                                                                                                                                                                                                                                                                                                             not null,
    metro          text[]       default ARRAY ['1'::text, '2'::text, '3'::text, '4'::text, '5'::text, '6'::text, '7'::text, '8'::text, '9'::text, '10'::text, '11'::text, '12'::text, '13'::text, '14'::text, '15'::text, '16'::text, '17'::text, '18'::text, '19'::text, '20'::text, '21'::text, '22'::text, '23'::text, '24'::text, '25'::text, '26'::text, '27'::text, '28'::text, '29'::text, '30'::text, '31'::text, '32'::text, '33'::text, '34'::text, '35'::text] not null,
    jk_ids         integer[]    default ARRAY []::integer[]                                                                                                                                                                                                                                                                                                                                                                                                               not null,
    silence        text         default ''::character varying                                                                                                                                                                                                                                                                                                                                                                                                             not null
);

create unique index filters_id_uindex
    on filters (id);

create unique index filters_user_id_uindex
    on filters (user_id);

create table metro
(
    id             uuid                     default uuid_generate_v4()    not null
        constraint metro_pk
            primary key,
    created_at     timestamp with time zone default CURRENT_TIMESTAMP     not null,
    metro          varchar(255)             default ''::character varying not null,
    transport_type varchar(255)             default ''::character varying not null,
    line_color     varchar(255)             default ''::character varying not null,
    time           integer                  default 0                     not null,
    line_id        varchar(255)             default ''::character varying not null,
    is_default     boolean                  default false                 not null,
    cian_id        integer                                                not null
        constraint metro_adverts_id_fk
            references adverts
);

create unique index metro_id_uindex
    on metro (id);


create table metro_directory
(
    line_name varchar(255) default ''::character varying not null,
    line_id   varchar(255) default ''::character varying not null
);



create table viewed
(
    id      uuid default uuid_generate_v4() not null
        constraint viewed_pk
            primary key,
    cian_id integer                         not null,
    user_id integer                         not null
);

create unique index viewed_id_uindex
    on viewed (id);

