# SQL Best Practices. Part 2

Databases are complex beasts, and that is inherently necessary. They are stateful, handle concurrency, handle persistence to disk, basically all the hard problems in programming. Additionally, they tend to have loose ownership, a big central source of truth that different applications and users consume and need to change continuously.

Thus, it is important to be specially careful about how we deal with them. Here are some good practices that I’ve gathered through the years. You can also check the first part in the previous [article](https://medium.com/@furstenheim/sql-best-practices-part-1-17f62d1b0f40).


## Field for created at and updated at
Nowadays, it is standard to use a Version Control System (say git) when we write code. This allows us to audit the changes in code and find when we introduced a bug. It also helps us understand in what circumstances we introduced certain code. This is not the case with databases, we only get to see the current snapshot of the data, we have no audit on how it got there.

A way to have control of the evolution of the records is to have an [audit table](https://dba.stackexchange.com/questions/15186/what-is-an-audit-table). Every time we perform an update or insert we save the modification. There is of course a huge impact on disk size. Another solution, would even be to use [git as a database](https://gitrows.com/).

However, there is a poor man's solution that is enough for most of the use cases, `created_at` and `updated_at`. We add two timestamps that we modify only when inserting or updating. We delegate to the database so there is no need to add it in any later application that accesses the database.

```
-- MySQL syntax
CREATE TABLE application_user (
    id_user bigserial primary key,
    ...
    created_at timestamp not null default CURRENT_TIMESTAMP
    updated_at timestamp null ON UPDATE CURRENT_TIMESTAMP
) COMMENT IS '...'
```

## Track database migrations
There is another side of not using VCS for the database. It's easier to have unsynchronized environments, for example, an index that exists in beta and production but not in gamma. It makes it hard to know when a change to the schema was introduced. Also, it misses code reviews, with all the goodies they bring: knowledge sharing, enforcing good practices... All of these can be addressed with database migrations.

In a repository hold all the database changes that are going to be executed, sorted by alphabetical order. Say 0001-start-database-schema.sql, 0002-include-new-column-for-user-favourite-application.sql and so on. There doesn't even need to be executed in a pipeline, it's ok to execute it manually, but it already gives you an audit of all changes that affect the database. Although it might seem that it only works for new databases, it can be used also for existing database, you just need to dump the schema without data and save it as first migration.

Once you have this in place, there are a couple of goodies that you can easily get from it. If you are using docker for development, you can make it read from that folder, thus ensuring that the database is updated and the tests run. That way, you can have several systems running their unit tests against the modified database with very low effort.
In case you are using Postgres, it's even nicer. Postgres has [transactionality for DDLs](https://julien.danjou.info/why-you-should-care-that-your-sql-ddl-is-transactional/), so it is possible to include these as part of the deployment pipeline and have them execute automatically, thus reducing access to databases. With other databases like MySQL it is not super recommendable because it would be possible to execute half a migration and have it fail in the middle, which makes it hard or impossible to record if a migration has passed or not.
Additionally, because you have a code review, you might be able to run a linter on the changes and ensure that they comply.


## Only use ON DELETE CASCADE when there is strong ownership
Foreign keys are the daily bread in database design. The obvious advantage is that they do not allow to insert wrong data into the database, but a second advantage is that they guide the user on the meaning of fields. Not all foreign keys are brewed alike, there are two inherently different kinds. Foreign keys that signal ownership and foreign keys that signal dependency.

Suppose the following scheme:


```

CREATE TABLE app_user_type (
    id_app_user_type bigserial primary key
    ...
)

CREATE TABLE app_user (
  id_user bigserial primary key
  id_app_user_type bigint references app_user_type (id_app_user_type) ON DELETE DO NOTHING
  ...
)

CREATE TABLE app_permission (
  id_permission bigserial primary key
  ...
)

CREATE TABLE app_user_has_app_permission (
  id_user bigint references app_user (id_user) ON DELETE CASCADE
  id_permission bigint references app_permission (id_permission) ON DELETE DO NOTHING
)
```

The foreign key between app_user_has_app_permission and app_user is there to signal ownership. Those permissions belong to an app user. If we delete an app user we want the whole list of permissions to be gone.

On the other hand, the foreign key from app_user_has_app_permission to app_permission is not of ownership but dependency. We want to understand what every permission mean. We wouldn't want to delete a permission if there are users with it, since most probably means that we would be removing functionality from them.

An even clearer example would be the foreign key from app_user to app_user_type. Users **do not belong** to a user type. It would be plain wrong to lose users if we delete a user type. We should get an error saying that there are still users there.

## JSON for non structured, non queried data
SQL is all about structure, your database schema dictates your access patterns. One common practice is to try to reflect all the structure of your data into the database. However, that is really not necessary, you only need to go as far as your query requirements and leave the rest as raw data.

To me, a very clear example is user configuration for a SPA (single page application). It might easily grow to be a middle size but very complex json, the ownership is fully on the frontend code with the backend probably doing schema validation. There is no need for the database to reflect that structure, when in reality the only thing that we want is to do the following query `SELECT uc.main_application_settings from user_configuration uc where uc.id_user = :id_user`, and we can send it all in one payload because it is only medium size. That way, a very reasonable schema would be the following:

```
CREATE TABLE user_configuration (
    id_user bigint primary key references application_user(id_user),
    user_configuration json
)
```

In case your database does not support json, you might even do a text field. In the end you are just writing in bulk and reading in bulk.

## Don't use mock databases for unit testing
Application development against any database follows the Pareto rule. The first 80% is relatively easy, setting up a schema that reflects your query patterns, writing the queries and so on, you might probably not use anything fancy for all this. However, the complexity is always on the nitty-gritty details. You are having a deadlock, do you need to change the concurrency control? Your query is slow, do you need to force an index or maybe write a CTE? You have to access several records for the same row, does your database allow window clauses? And so on.

Databases provide us with small knobs that we can touch to get exactly what we want. And each of these are slightly different in different databases. The problem with using a mock database is that your queries now need to run against two different databases, the mock database and the real database and each behave slightly different, specially for these tiny details. In the end this is completely limiting your toolbox of solutions.

Not only you are limiting your toolbox, there are some issues that only occur in some databases. Maybe the mock database does not access the filesystem, so it doesn't get some sort of race condition. Or maybe there is a bug in your current database that cannot show in the mock database. If you are using a mock database for unit testing, that means that developers need to use integration tests to reproduce some failures, which is way more complex and sometimes impossible, because you are not on the same process.

Ok, so no mock databases, what do you do? Easy, we are in 2021 and containers are easy to set up. Create a container with exactly the same version that you have for production and develop against that. It won't have data and the connection is in localhost, so all tests will run fast, and they will reflect exactly the behavior that you have in production. Moreover, if you are using database migrations you can set up your docker from read from them, thus ensuring that your application can execute against the latest database changes.

![](./test-database.png)

## Avoid stored procedures
As I mentioned in the introduction, databases tend to have loose ownership. Stored procedures mean having logic in the database, which not a clear owner. In general, it's better to have this logic in a library as code and import it as needed. It will make it easier to roll out changes, and make it clear whether the code can be deprecated.

That was all. Happy coding!

