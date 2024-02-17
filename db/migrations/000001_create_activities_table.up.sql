BEGIN;

CREATE TABLE activities
(
    id              bigserial       CONSTRAINT activities_pk PRIMARY KEY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    content_uuid    uuid            NOT NULL,

    name            varchar         NOT NULL,
    description     text            NULL,
    content_type    text            NOT NULL,
    taxonomy        text            NULL,

    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX activities_id_uindex
    ON activities (id);
CREATE UNIQUE INDEX activities_uuid_uindex
    ON activities (uuid);

COMMIT;
