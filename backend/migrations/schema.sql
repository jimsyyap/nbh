-- Users table
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  role VARCHAR(50) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Posts table
CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  slug VARCHAR(255) UNIQUE NOT NULL,
  content TEXT NOT NULL,
  excerpt TEXT,
  author_id INTEGER REFERENCES users(id),
  status VARCHAR(50) NOT NULL,
  featured BOOLEAN DEFAULT false,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  published_at TIMESTAMP WITH TIME ZONE
);

-- Categories table
CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  slug VARCHAR(100) UNIQUE NOT NULL,
  description TEXT
);

-- Post categories (many-to-many)
CREATE TABLE post_categories (
  post_id INTEGER REFERENCES posts(id),
  category_id INTEGER REFERENCES categories(id),
  PRIMARY KEY (post_id, category_id)
);

-- Comments table
CREATE TABLE comments (
  id SERIAL PRIMARY KEY,
  post_id INTEGER REFERENCES posts(id),
  user_id INTEGER REFERENCES users(id),
  content TEXT NOT NULL,
  status VARCHAR(50) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Media table
CREATE TABLE media (
  id SERIAL PRIMARY KEY,
  filename VARCHAR(255) NOT NULL,
  path VARCHAR(255) NOT NULL,
  type VARCHAR(100) NOT NULL,
  size INTEGER NOT NULL,
  alt_text VARCHAR(255),
  user_id INTEGER REFERENCES users(id),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
