CREATE TABLE list_users (
    list_id UUID NOT NULL REFERENCES lists(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role TEXT NOT NULL CHECK (role IN ('editor', 'viewer')),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (list_id, user_id)
);

CREATE INDEX idx_list_users_user_id ON list_users(user_id);
CREATE INDEX idx_list_users_list_id ON list_users(list_id);