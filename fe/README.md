# zydukas-rankings — frontend

React + TypeScript + Vite. Google sign-in through Firebase, the ID token goes
to the Go API in `../be`.

## Running

```bash
cp .env.example .env    # fill in the Firebase web config
npm install
npm run dev             # http://localhost:5173
```

The Go API has no CORS middleware, so Vite proxies `/api` to
`http://localhost:8080` and strips the prefix (`vite.config.ts`). Override the
target with `BACKEND_URL`.

## API used

Only what `../be/app/app.go` actually routes:

| Call | Endpoint | Auth |
| --- | --- | --- |
| `APIs/userAPI.ts` → `createUser` | `POST /users` | yes |
| `APIs/userAPI.ts` → `getUser` | `GET /users` | yes |
| `APIs/userAPI.ts` → `updateUsername` | `PATCH /users/:new_username` | yes |
| `APIs/userAPI.ts` → `deleteUser` | `DELETE /users` | yes |
| `APIs/postsAPI.ts` → `uploadPost` | `POST /posts/upload` | no (route has no middleware) |
| `APIs/postsAPI.ts` → `getFeed` | `GET /posts` | **not routed yet** |

Requests go through one plain axios instance (`APIs/client.ts`) with two
interceptors: the request one adds the Firebase `Authorization` header, the
response one turns every failure into an `ApiError` (status + message). Nothing
else about axios is used — no baseURL, no transforms, no retries.

Payloads are `{"data": ...}` on success, `{"message": "..."}` for
acknowledgements and `{"error": "..."}` (sometimes nested) on failure;
`readMessage` in `Errors/exceptionHandler.ts` unwraps all three for the
interceptor.

The feed page is built, the endpoint behind it is not. `POST /posts/upload`
writes the file to `be/output/` and stores nothing, and `be/app/app.go` routes
no `GET /posts`, so `getFeed` treats 404 and 500 as an empty feed (the same way
`getUser` treats a missing row) and the page shows its empty state instead of a
toast on every load. Add the route plus a `models.Post` with `id`, `username`,
`image_url`, `created_at` and `caption` (see `PostDTO`) and the feed fills
itself.

The upload sends the caption as a `caption` field in the same multipart body as
the file, and the feed card renders it when the row has one. `CreatePost` only
reads `file`, so the caption is dropped until that `models.Post` exists —
`PostDTO.caption` is optional for the same reason.

## Layout

```
src/
  APIs/       client.ts (axios + interceptors) + one function per endpoint
  Auth/       firebase init, sign-in, useAuthUser
  Errors/     ApiError, the message unwrapper, the toast list
  Store/      redux: global errors only
  Utils/      DTOs, query keys
  TopBar/     the chrome
  Pages/      FeedPage.tsx — the app shell (login / register / app gate) and
              the feed itself; Upload, Profile, Register
```

Server data goes through react-query (`USER_INFO_QUERY_KEY`), client state
through redux. Failed requests raise one toast from the query/mutation cache in
`main.tsx`; components don't handle errors themselves.
