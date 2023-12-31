# Golang Fiber Full-stack Web App

### Simple web application full-stack to play Trivia + API with Supabase as database + CRUD, made with Go Fiber, HTML template rendering and </> htmx & _hyperscript.

---

### Setup:

Besides the obvious prerequisite of having Go!, you must have Air installed for hot reloading when editing code.

On the other hand, you must have an account in Supabase and within a project you must create the "facts" table using the following SQL statement:

```
create table
  public.facts (
    id bigint generated by default as identity,
    created_at timestamp with time zone not null default now(),
    question text not null,
    answer text not null,
    constraint facts_pkey primary key (id),
    constraint facts_id_key unique (id)
  ) tablespace pg_default;
```

The access credentials to your Supabase project must appear in an .env file:

```
SUPABASE_URL=xxxx
SUPABASE_KEY=xxxx
```

---

#### Start the App in development mode:

```
$ air  # Ctrl + C to stop the app
```

#### Compile for production:

```
$ go build -ldflags="-s -w" -o ./cmd/main ./cmd/main.go  # ./cmd/main to run the app
```
