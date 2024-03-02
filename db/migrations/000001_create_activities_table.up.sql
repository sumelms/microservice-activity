BEGIN;

CREATE TABLE activities
(
    id              integer         PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    content_uuid    uuid            NOT NULL, -- Main content of the activity. This content has a specific type
                                              -- (video, pdf, questionnaire, slide, etc.), each type of content
                                              -- is implemented in a dedicated Micro Service.
    name            varchar         NOT NULL,
    description     text            NULL,
    content_type    text            NOT NULL,
    taxonomy        text            NULL,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX activities_uuid_uindex
    ON activities (uuid);

COMMIT;
