CREATE TABLE IF NOT EXISTS lrc (
    tweet_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    likes INT DEFAULT 0 NOT NULL,
    comments INT DEFAULT 0 NOT NULL,
    retweet INT DEFAULT 0 NOT NULL
);


create table if not exists liked(
    tweet_id int,
    username varchar(255),
    liked_user varchar(255)
);

create table if not exists commented(
    tweet_id int,
    username varchar(255),
    commented_user varchar(255),
    comment varchar(500)
);

create table if not exists retweet(
    tweet_id int,
    username varchar(255),
    retweet_user varchar(255)
);