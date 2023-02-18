CREATE TABLE IF NOT EXISTS word (
                                        word_id TEXT,
                                        word TEXT,
                                        phonetic TEXT,
                                        definition TEXT,
                                        difficulty TEXT,
                                        last_used BIGINT,
                                        num_repetitions INT,
                                        create_time BIGINT NOT NULL,
                                        update_time BIGINT NOT NULL,
                                        PRIMARY KEY(word_id)
) WITHOUT ROWID;

CREATE TABLE IF NOT EXISTS phrase(
                                     phrase_id TEXT,
                                     phrase TEXT,
                                     create_time BIGINT NOT NULL,
                                     update_time BIGINT NOT NULL,
                                     PRIMARY KEY(phrase_id)
) WITHOUT ROWID;




CREATE TABLE IF NOT EXISTS word_phrase(
                                          word_phrase_id TEXT,
                                          word_id TEXT,
                                          phrase_id TEXT,
                                          create_time BIGINT NOT NULL,
                                          update_time BIGINT NOT NULL,
                                          PRIMARY KEY(word_phrase_id)
) WITHOUT ROWID;
