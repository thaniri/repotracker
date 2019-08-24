CREATE DATABASE IF NOT EXISTS repotrackerapp CHARACTER SET utf8;

CREATE TABLE IF NOT EXISTS repotrackerapp.rpUsers
  (
     id           INT UNSIGNED NOT NULL PRIMARY KEY auto_increment,
     email        VARCHAR(255) NOT NULL,
     passwordHash VARCHAR(255) NOT NULL,
     registered   BOOLEAN      NOT NULL DEFAULT 0
  );

CREATE TABLE IF NOT EXISTS repotrackerapp.githubRepos
  (
     id      INT UNSIGNED NOT NULL PRIMARY KEY auto_increment,
     repoUrl VARCHAR(255) NOT NULL
  );

CREATE TABLE IF NOT EXISTS repotrackerapp.rpUsersFollowingGithubRepos
  (
     id           INT UNSIGNED NOT NULL PRIMARY KEY auto_increment,
     rpUserId     INT UNSIGNED NOT NULL,
     githubRepoId INT UNSIGNED NOT NULL,
     CONSTRAINT fk_rpuserid FOREIGN KEY (rpUserId) REFERENCES rpUsers(id),
     CONSTRAINT fk_githubrepoid FOREIGN KEY (githubRepoId) REFERENCES
     githubRepos(id)
  ); 

