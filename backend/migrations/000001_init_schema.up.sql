CREATE TABLE IF NOT EXISTS games (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    board TEXT[] DEFAULT ARRAY['', '', '', '', '', '', '', '', ''],
    current_turn VARCHAR(1) DEFAULT 'X',
    status VARCHAR(20) DEFAULT 'waiting',
    player_x_id INTEGER REFERENCES users(id),
    player_o_id INTEGER REFERENCES users(id),
    winner VARCHAR(1),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    finished_at TIMESTAMP WITH TIME ZONE
);
