Mini Twitter Project Overview

# This is a Mini Twitter Project consisting of five microservices.

# Technologies Used

-   Microservices Architecture

-   PostgreSQL

-   Apache Kafka

-   Redis

-   SQL Builder

-   HTTPS/HTTP

-   JWT Tokens

-   Docker

-   CI/CD



### 1. User Service

This service handles user-related functionalities, including:

* Authorization: User registration, login, logout, and account deletion.

* Profile Management: Viewing and editing the user profile.

* Followers and Following:

* View your followers.

* See who you are following.

* Follow or unfollow users.

* Search for other users.

# Note: All user GET requests are performed using the username, which is parsed from the JWT token. Users do not need to manually provide their username.

### 2. Tweet Service

This service allows users to manage their tweets, including:

* Creating a new tweet.

* Updating an existing tweet.

* Deleting a tweet.

* Viewing a specific tweet.

* Viewing all tweets created by the user.

### 3. LRC (Like, Retweet, Comment) Service

This service handles interactions with tweets, including:

# Like:

Users can like other users' tweets.

Notifications are sent to the tweet owner when their tweet is liked.

# Comment:

Users can comment on tweets.

Notifications are sent to the tweet owner when comments are made.

# Retweet:

Users can retweet other tweets.

The retweeted tweet is added to the user's own tweets list.

Additionally, users can:

View who liked their tweet.

View all comments on their tweet.

View the list of users who retweeted their tweet.

Note: To retrieve this information, users must provide the tweet ID when making API requests.

### 4. Notification Service

This service handles real-time notifications.

To connect:

Send a request to the WebSocket route localhost:8083/ws.

Include a header with the format: username:{your_username}.

Users can then receive notifications related to their tweets based on their username.

### 5. API Gateway

The API Gateway serves as the central point for:

Forwarding requests to the appropriate microservices.

Returning responses to the user.

Features:

Built on the HTTP request-response system, I know how to work with Gin framework. However I feel comfortable to use HTTP .

Uses JWT tokens for authorization.

Token Requirements:

No token is needed for registration or email verification.

A token is required for all other operations.

After registration, a confirmation code is sent to the user's email. Once verified, the user receives a JWT token.

### Overall Project Summary

This project was developed within a tight deadline, and I did my best to implement the core functionalities.

Future updates could further enhance the project, especially improving the follow, unfollow, and retweet features to match the performance of real-world