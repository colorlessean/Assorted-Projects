# Database Schema
This document will detail the basic database schema. Detailing tables, attributes and relations.

## User
User table will have the following attributes
* username (key, unique)
* email (unique)
* bio

## Follow
Follow table will show who is following who. It is a relation table with the attributes:
* user (foreign key User.username)
* follower (foreign key User.username)
Can get the users followers list and following list from this table.

## Post
Post table will contain all the posts by users. Has the following attributes.
* postID (key, db generated)
* postType (should be an enum of default, repost, and reply)
* username
* content

## PostRelation
PostRelation is a relation table that contains the parent to child relationship for reposting and replying. Has the attributes:
* parent (foreign key Post.postID)
* child (foreign key Post.postID)

## Likes
Likes table will show who likes what post. Relational table with the following attributes:
* post (foreign key Post.postID)
* username (foreign key Post.post)

This should be a sufficient list of tables to get started with.