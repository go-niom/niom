package blog

// categories table
var PGCreateCategoriesTable = `
CREATE TABLE categories (
  cat_id INTEGER NOT NULL,
  cat_title VARCHAR(235) NOT NULL
);
`

// comment table
var PGCreateCommentTable = `
CREATE TABLE comments (
  comm_id INTEGER NOT NULL,
  comm_author VARCHAR(255) NOT NULL,
  comm_content TEXT NOT NULL,
  comm_date DATE NOT NULL,
  comm_status VARCHAR(255) NOT NULL,
  comm_post_id INTEGER NOT NULL,
  comm_email VARCHAR(255) NOT NULL
);
`

// post table
var PGCreatePostTable = `
CREATE TABLE posts (
  post_id INTEGER NOT NULL,
  post_category_id INTEGER NOT NULL,
  post_title VARCHAR(255) NOT NULL,
  post_author VARCHAR(255) NOT NULL,
  post_date DATE NOT NULL,
  post_image TEXT NOT NULL,
  post_content TEXT NOT NULL,
  post_tags VARCHAR(255) NOT NULL,
  post_comment_count INTEGER NOT NULL,
  post_status VARCHAR(255) NOT NULL DEFAULT 'draft'
);
`

// users table
var PGCreateUsersTable = `
CREATE TABLE users (
  user_id INTEGER NOT NULL,
  username VARCHAR(225) NOT NULL,
  user_password VARCHAR(225) NOT NULL,
  user_firstname VARCHAR(225) NOT NULL,
  user_lastname VARCHAR(225) NOT NULL,
  user_email VARCHAR(225) NOT NULL,
  user_image TEXT NOT NULL,
  user_role VARCHAR(225) NOT NULL,
  rant_solt VARCHAR(225) NOT NULL
);
`

var PGDropCategoriesTable = `DROP TABLE IF EXISTS categories;`

var PGDropCommentTable = `DROP TABLE IF EXISTS comments;`

var PGDropPostTable = `DROP TABLE IF EXISTS posts;`

var PGDropUsersTable = `DROP TABLE IF EXISTS users;`

var PGBlogCategoriesInsert = `INSERT INTO categories (cat_id, cat_title) VALUES
        (1, 'Java'),
        (2, 'Python'),
        (13, 'OOP'),
        (15, 'c');`

var PGBlogCommentInsert = `INSERT INTO comments (comm_id, comm_author, comm_content, comm_date, comm_status, comm_post_id, comm_email) VALUES
        (6, 'Suraj kumar jha', 'ddddd', '2020-02-09', 'Unapproved', 1, 'user@admin.com'),
        (7, 'Suraj kumar jha', 'vvvvvhjv', '2020-02-09', 'Unapproved', 1, 'user@admin.com'),
        (8, 'Suraj kumar jha', 'kjbdjgsfkjhgdfsl', '2020-02-09', 'Approved', 1, 'user@admin.com'),
        (9, 'ncjxg.kjxf', 'bbbb', '2020-02-09', 'Approved', 5, 'user@admin.com'),
        (10, 'Suraj kumar jha', 'bbb', '2020-02-09', 'Approved', 5, 'event@admin.com'),
        (12, '', '', '2020-03-05', 'draft', 9, ''),
        (13, 'Suraj kumar jha', '', '2020-03-05', 'draft', 9, ''),
        (14, 'Suraj kumar jha', '', '2020-03-05', 'draft', 9, ''),
        (15, 'nfuyhjfkj,', 'mngyiulgk;ol/', '2020-03-07', 'Approved', 1, 'kumarjhasuraj6@gmail.com');`

var PGBlogPostInsert = `INSERT INTO posts (post_id, post_category_id, post_title, post_author, post_date, post_image, post_content, post_tags, post_comment_count, post_status) VALUES
        (1, 1, 'java', 'dfghjkl', '2020-03-07', 'java.png', '<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Reiciendis aliquid atque, nulla? Quos cum ex quis soluta, a laboriosam. Dicta expedita corporis animi vero voluptate voluptatibus possimus, veniam magni quis!</p>\r\n', 'AI,machine', 12, 'published'),
        (5, 1, 'py', 'xcvbnm,.dsfghjk', '2020-03-07', 'download.jpg', '<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Reiciendis aliquid atque, nulla? Quos cum ex quis soluta, a laboriosam. Dicta expedita corporis animi vero voluptate voluptatibus possimus, veniam magni quis!</p>\r\n', 'AI,machine', 18, 'published'),
        (7, 1, 'Ai', 'dfghjkl', '2020-03-07', 'download.jpg', '<p>nbchg</p>\r\n', 'AI,machine,c++', 4, 'published'),
        (9, 1, 'fhsgladfsli', 'kjsbgvidhgi', '2020-02-22', 'java.png', 'kjzshkbrydiuethvbiusvyherakbtuheritbuearibthoer', 'AI,machine,c++', 19, 'published');`

var PGBlogUserInsert = `INSERT INTO users (user_id, username, user_password, user_firstname, user_lastname, user_email, user_image, user_role, rant_solt) VALUES
        (20, 'opt6@lpu.com', '123', 'Suraj', 'jhaguukghv', 'event123@admin.com', 'java.png', 'admin', ''),
        (21, 'suraj', '321', 'Suraj', 'jhaguukghv', 'event123@admin.com', 'java.png', 'admin', ''),
        (22, 'qqq', 'qqq', 'Suraj', 'jha', 'kumarjhasuraj6@gmail.com', '', 'admin', '');`

var PGBlogCategoriesDelete = `DELETE FROM categories WHERE cat_id IN (1, 2, 13, 15);`

var PGBlogCommentDelete = `DELETE FROM comments WHERE comm_id IN (6, 7, 8, 9, 10, 12, 13, 14, 15);`

var PGBlogPostDelete = `DELETE FROM posts WHERE post_id IN (1, 5, 7, 9);`

var PGBlogUserDelete = `DELETE FROM users WHERE user_id IN (20, 21, 22);`

var PGBlogAlter = `ALTER TABLE categories ADD PRIMARY KEY (cat_id);

ALTER TABLE comments ADD PRIMARY KEY (comm_id);

ALTER TABLE posts ADD PRIMARY KEY (post_id);

ALTER TABLE users ADD PRIMARY KEY (user_id);

ALTER TABLE categories ALTER COLUMN cat_id SET NOT NULL;
ALTER TABLE categories ALTER COLUMN cat_id ADD GENERATED ALWAYS AS IDENTITY;

ALTER TABLE comments ALTER COLUMN comm_id SET NOT NULL;
ALTER TABLE comments ALTER COLUMN comm_id ADD GENERATED ALWAYS AS IDENTITY;

ALTER TABLE posts ALTER COLUMN post_id SET NOT NULL;
ALTER TABLE posts ALTER COLUMN post_id ADD GENERATED ALWAYS AS IDENTITY;

ALTER TABLE users ALTER COLUMN user_id SET NOT NULL;
ALTER TABLE users ALTER COLUMN user_id ADD GENERATED ALWAYS AS IDENTITY;`

var PGBlogOppositeAlter = `ALTER TABLE categories DROP CONSTRAINT categories_pkey;

ALTER TABLE comments DROP CONSTRAINT comments_pkey;

ALTER TABLE posts DROP CONSTRAINT posts_pkey;

ALTER TABLE users DROP CONSTRAINT users_pkey;

ALTER TABLE categories ALTER COLUMN cat_id DROP IDENTITY;
ALTER TABLE categories ALTER COLUMN cat_id DROP NOT NULL;

ALTER TABLE comments ALTER COLUMN comm_id DROP IDENTITY;
ALTER TABLE comments ALTER COLUMN comm_id DROP NOT NULL;

ALTER TABLE posts ALTER COLUMN post_id DROP IDENTITY;
ALTER TABLE posts ALTER COLUMN post_id DROP NOT NULL;

ALTER TABLE users ALTER COLUMN user_id DROP IDENTITY;
ALTER TABLE users ALTER COLUMN user_id DROP NOT NULL;`
