# OneLink

No matter how many links — share one.

## The Problem

Sharing multiple links is messy. You copy a bunch of URLs, dump them in a chat, maybe pin the message if you need them later. It's sloppy and slow.

You should be able to go to one place, see all the links, and click what you want. What better place than a link itself?

## How It Works

1. Go to OneLink
2. Paste your links (add titles, notes, whatever)
3. Get one short link back
4. Share it
5. Anyone with that link sees everything — clean, fast, no friction

No accounts. No logins. No sign-up walls. Just paste and share.

## Tech Stack

**Frontend:** Vite, React, TailwindCSS, React Hook Form + Zod, TanStack Query

**Backend:** Go, Chi, sqlx, PostgreSQL, slog

**Infra:** Docker, Fly.io

**API Contract:** OpenAPI spec

## Inspiration

If [ilovepdf.com](https://www.ilovepdf.com) can let you merge PDFs without logging in, you should be able to bundle links the same way.