-- Drop existing users table if it exists
DROP TABLE IF EXISTS users;

-- Recreate users table to match models.User
CREATE TABLE users (
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

-- Insert staff user
-- UserTypeStaff = 1
-- bcrypt hash corresponds to password: "poc234"
INSERT INTO users (username, password)
VALUES (
    'staff',
    'poc123'
);
