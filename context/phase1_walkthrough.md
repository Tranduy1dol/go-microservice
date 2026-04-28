# Phase 1 Walkthrough — Foundation

## Step Overview

| # | Task | Status |
|---|------|--------|
| 1 | Clean old code + restructure directories | ⬜ |
| 2 | Update `go.mod` with new dependencies | ⬜ |
| 3 | Docker Compose (MongoDB + Redis) | ⬜ |
| 4 | Config layer (`config/config.go`) | ⬜ |
| 5 | Domain models (Word, Question, Paragraph, Grammar, Test, User, SRS) | ⬜ |
| 6 | Port interfaces | ⬜ |
| 7 | MongoDB adapter — connect + Word repo | ⬜ |
| 8 | MongoDB adapter — remaining repos | ⬜ |
| 9 | Use cases (Lookup, basic CRUD) | ⬜ |
| 10 | Google OAuth adapter | ⬜ |
| 11 | HTTP router + middleware + handlers | ⬜ |
| 12 | Admin endpoints | ⬜ |
| 13 | JMdict importer CLI | ⬜ |
| 14 | Verify: `docker-compose up` + import + query | ⬜ |

---

## Step 1: Clean Old Code + Restructure

**Goal:** Remove the old social-media code, create the hexagonal layout.

```bash
# Remove old source files (keep .git, go.mod, .air.toml, .gitignore)
rm -rf cmd/ internal/ docs/ scripts/

# Create new directory structure
mkdir -p cmd/api cmd/importer
mkdir -p internal/domain internal/port internal/usecase
mkdir -p internal/adapter/mongo internal/adapter/redis
mkdir -p internal/adapter/searchgrpc internal/adapter/translation
mkdir -p internal/adapter/auth internal/adapter/jmdict
mkdir -p api/handler api/middleware api/dto
mkdir -p config proto/search
```

---

## Step 2: Update `go.mod`

Reset dependencies for the new stack:

```bash
# Reset go.mod
go mod init github.com/Tranduy1dol/go-microservice

# Add core dependencies
go get github.com/gin-gonic/gin@latest
go get go.mongodb.org/mongo-driver/v2@latest
go get github.com/redis/go-redis/v9@latest
go get github.com/spf13/viper@latest
go get go.uber.org/zap@latest
go get github.com/go-playground/validator/v10@latest
go get github.com/spf13/cobra@latest
go get golang.org/x/oauth2@latest
go get github.com/golang-jwt/jwt/v5@latest
go get github.com/stretchr/testify@latest
go get github.com/swaggo/swag/cmd/swag@latest
go get github.com/swaggo/gin-swagger@latest
go get github.com/swaggo/files@latest
```

---

## Step 3: Docker Compose

Replace `docker-compose.yml` with MongoDB + Redis:

```yaml
services:
  mongo:
    image: mongo:7
    container_name: nihongo_mongo
    ports:
      - "27017:27017"
    environment:
      MONGO_INITDB_ROOT_USERNAME: admin
      MONGO_INITDB_ROOT_PASSWORD: secret
    volumes:
      - mongo_data:/data/db

  redis:
    image: redis:7-alpine
    container_name: nihongo_redis
    ports:
      - "6379:6379"

volumes:
  mongo_data:
```

Start it: `docker-compose up -d`

---

## Step 4: Config (`config/config.go`)

Use Viper to load config from env/file:

```go
package config

type Config struct {
    Server   ServerConfig
    MongoDB  MongoConfig
    Redis    RedisConfig
    OAuth    OAuthConfig
}

type ServerConfig struct {
    Port string
    Env  string
}

type MongoConfig struct {
    URI      string
    Database string
}

type RedisConfig struct {
    Addr     string
    Password string
    DB       int
}

type OAuthConfig struct {
    GoogleClientID     string
    GoogleClientSecret string
    RedirectURL        string
    JWTSecret          string
}
```

Load with `viper.AutomaticEnv()` + a `.env` file.

---

## Step 5: Domain Models

Create these files in `internal/domain/`:

1. **`word.go`** — Word, Kanji, Reading, Sense, Gloss
2. **`question.go`** — Question, QuestionType enum
3. **`paragraph.go`** — Paragraph with embedded questions
4. **`grammar.go`** — Grammar, GrammarExample
5. **`test.go`** — Test, TestPart, TestSection, TestResult
6. **`user.go`** — User (Google OAuth fields), StudyProgress
7. **`srs.go`** — SRSCard, SM-2 algorithm function

> [!TIP]
> Domain models should have **zero imports** from external packages (except `time`). Use `bson` struct tags for MongoDB serialization. The domain layer is pure business logic.

---

## Step 6: Port Interfaces

Create these in `internal/port/`:

1. **`dictionary.go`** — `DictionaryRepository` interface (CRUD + search by kanji/reading/JLPT + random)
2. **`question.go`** — `QuestionRepository` interface
3. **`paragraph.go`** — `ParagraphRepository` interface
4. **`grammar.go`** — `GrammarRepository` interface
5. **`test.go`** — `TestGenerator` + `TestRepository` interfaces
6. **`search.go`** — `SearchEngine` interface (placeholder for Phase 2)
7. **`user.go`** — `UserRepository` interface
8. **`translation.go`** — `Translator` interface (placeholder for Phase 3)

---

## Step 7: MongoDB Adapter — Word Repo

First file in `internal/adapter/mongo/`:

1. **`client.go`** — MongoDB connection setup, returns `*mongo.Database`
2. **`word_repo.go`** — Implements `port.DictionaryRepository`
   - Collection: `words`
   - Key methods: `GetByID`, `GetByKanji`, `Create`, `BulkCreate`, `GetRandom`
   - Use MongoDB aggregation `$sample` for random word selection

---

## Step 8: Remaining MongoDB Repos

- **`question_repo.go`** — Collection: `questions`, sampling by JLPT + section
- **`paragraph_repo.go`** — Collection: `paragraphs`
- **`grammar_repo.go`** — Collection: `grammar`
- **`user_repo.go`** — Collection: `users`, upsert on OAuth login

---

## Step 9: Use Cases

In `internal/usecase/`:

1. **`lookup.go`** — LookupService: word/grammar lookup orchestration
2. **`test_generator.go`** — Assembles JLPT mock tests from content pools

---

## Step 10: Google OAuth

In `internal/adapter/auth/`:

1. **`google_oauth.go`** — OAuth2 flow (redirect → callback → get user info)
2. **`jwt.go`** — Issue + validate JWT tokens for authenticated sessions

In `api/middleware/`:
1. **`auth.go`** — Extract JWT from `Authorization` header, inject user into context
2. **`admin.go`** — Check user role == "admin"

---

## Step 11: HTTP Router + Handlers

1. **`api/router.go`** — Gin router with all route groups
2. **`api/handler/word_handler.go`** — GET word, search, browse JLPT
3. **`api/handler/grammar_handler.go`** — GET grammar
4. **`api/handler/test_handler.go`** — POST generate, POST submit
5. **`api/handler/user_handler.go`** — GET /me, GET /progress
6. **`cmd/api/main.go`** — Wire everything, start server

---

## Step 12: Admin Handlers

**`api/handler/admin_handler.go`** — CRUD for all content types behind admin middleware.

---

## Step 13: JMdict Importer

**`cmd/importer/main.go`** + **`internal/adapter/jmdict/parser.go`**

- Download JMdict_e.xml.gz from edrdg.org
- Parse XML → `[]domain.Word`
- Bulk insert into MongoDB

---

## Step 14: Verify

```bash
docker-compose up -d
go run ./cmd/importer/main.go --file=JMdict_e.xml
go run ./cmd/api/main.go
curl http://localhost:8080/api/v1/health
curl http://localhost:8080/api/v1/words/search?q=taberu
```
