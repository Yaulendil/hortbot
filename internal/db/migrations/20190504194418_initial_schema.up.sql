-- This file is subject to change until the first real deployment of the bot.
-- Do not rely on these schema migrations remaining consistent until this
-- message has been removed.

BEGIN;

CREATE TYPE access_level AS ENUM (
    'everyone',
    'subscriber',
    'moderator',
    'broadcaster',
    'admin'
);

CREATE TABLE channels (
    id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,

    user_id bigint NOT NULL UNIQUE,
    name text NOT NULL,
    bot_name text NOT NULL,
    active boolean NOT NULL,
    prefix text NOT NULL,
    bullet text,

    ignored text[] DEFAULT '{}' NOT NULL,
    custom_owners text[] DEFAULT '{}' NOT NULL,
    custom_mods text[] DEFAULT '{}' NOT NULL,
    custom_regulars text[] DEFAULT '{}' NOT NULL,

    cooldown int,
    last_command_at timestamptz NOT NULL,

    last_fm text NOT NULL,
    parse_youtube boolean NOT NULL,

    should_moderate boolean NOT NULL,
    display_warnings boolean NOT NULL,
    enable_warnings boolean NOT NULL,
    timeout_duration int NOT NULL,
    enable_filters boolean NOT NULL,

    filter_links boolean NOT NULL,
    permitted_links text[] DEFAULT '{}' NOT NULL,

    filter_caps boolean NOT NULL,
    filter_caps_min_chars int NOT NULL,
    filter_caps_percentage int NOT NULL,
    filter_caps_min_caps int NOT NULL,

    filter_emotes boolean NOT NULL,
    filter_emotes_max int NOT NULL,
    filter_emotes_single boolean NOT NULL,

    filter_symbols boolean NOT NULL,
    filter_symbols_percentage int NOT NULL,
    filter_symbols_min_symbols int NOT NULL,

    filter_me boolean NOT NULL,
    filter_max_length int NOT NULL,

    filter_banned_phrases boolean NOT NULL,
    filter_banned_phrases_patterns text[] DEFAULT '{}' NOT NULL,

    CHECK (filter_caps_percentage BETWEEN 0 and 100),
    CHECK (filter_symbols_percentage BETWEEN 0 and 100),
    CHECK (timeout_duration >= 0)
);

CREATE INDEX channels_user_id_idx on channels (user_id);

CREATE TABLE simple_commands (
    id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,

    channel_id bigint REFERENCES channels (id) NOT NULL,

    name text NOT NULL,
    message text NOT NULL,
    access_level access_level NOT NULL,
    count bigint NOT NULL,

    creator text NOT NULL,
    editor text NOT NULL,

    UNIQUE (channel_id, name)
);

CREATE INDEX simple_commands_channel_id_idx ON simple_commands (channel_id);
CREATE INDEX simple_commands_name_idx on simple_commands (name);

CREATE TABLE quotes (
    id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,

    channel_id bigint REFERENCES channels (id) NOT NULL,

    num int NOT NULL,
    quote text NOT NULL,

    creator text NOT NULL,
    editor text NOT NULL,

    UNIQUE (channel_id, num)
);

CREATE INDEX quotes_channel_id_idx ON quotes (channel_id);
CREATE INDEX quotes_num_idx on quotes (num);

CREATE TABLE repeated_commands (
    id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,

    channel_id bigint REFERENCES channels (id) NOT NULL,
    simple_command_id bigint REFERENCES simple_commands (id) NOT NULL UNIQUE,

    enabled boolean NOT NULL,
    delay int NOT NULL,
    message_diff bigint DEFAULT 1 NOT NULL,
    last_count bigint NOT NULL,

    creator text NOT NULL,
    editor text NOT NULL
);

CREATE INDEX repeated_commands_channel_id_idx ON repeated_commands (channel_id);
CREATE INDEX repeated_commands_simple_command_id_idx ON repeated_commands (simple_command_id);

CREATE TABLE scheduled_commands (
    id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,

    channel_id bigint REFERENCES channels (id) NOT NULL,
    simple_command_id bigint REFERENCES simple_commands (id) NOT NULL UNIQUE,

    enabled boolean NOT NULL,
    cron_expression text NOT NULL,
    message_diff bigint DEFAULT 1 NOT NULL,
    last_count bigint NOT NULL,

    creator text NOT NULL,
    editor text NOT NULL
);

CREATE INDEX scheduled_commands_channel_id_idx ON scheduled_commands (channel_id);
CREATE INDEX scheduled_commands_simple_command_id_idx ON scheduled_commands (simple_command_id);

CREATE TABLE autoreplies (
    id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,

    channel_id bigint REFERENCES channels (id) NOT NULL,

    num int NOT NULL,
    trigger text NOT NULL,
    orig_pattern text,
    response text NOT NULL,
    count int NOT NULL,

    creator text NOT NULL,
    editor text NOT NULL
);

CREATE INDEX autoreplies_channel_id_idx ON autoreplies (channel_id);
CREATE INDEX autoreplies_num_idx ON autoreplies (num);

COMMIT;
