# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project概要

Among Us進行補助bot — Go言語で書かれたDiscordボット。Among Usゲームのボイスチャンネルミュート管理を自動化します。ミーティングボイスチャンネル内のプレイヤーのミュートを一括切り替えし、ロビーとミーティングボイスチャンネル間を移動するプレイヤーを自動ミュート/ミュート解除します。

## 開発コマンド

### ローカル
```bash
docker-compose up          # starts bot + Redis
docker-compose up --build  # rebuild and start
```

[Air](https://github.com/air-verse/air) によるHotReload環境を作っています。

## 環境変数

```
cp .env.example .env
```
その後変数を埋めて下さい

| 変数 | もく亭 |
|---|---|
| `ACCESS_TOKEN` | Discord bot token |
| `CLIENT_ID` / `CLIENT_SECRET` | Discord bot credentials |
| `KVS_HOST` / `KVS_PORT` | Redis接続Host |
| `CONTROL_TEXT_CH` | mを受け入れるch名 |
| `LOBBY_VC` | ロビーVC名 |
| `MEETING_VC` | 会議VC名 |

## 構成

```
app/
├── main.go                           # Entry point: creates Discord session, registers handlers
├── configs/environment.go            # Global config loaded from env vars at init
├── cmd/
│   ├── usecases/
│   │   ├── message_handle/main.go    # Handles "m" command → bulk mute toggle
│   │   └── voice_handle/main.go      # Handles VC move events → auto mute/unmute
│   └── utils/discord_utils.go        # Shared Discord API helpers
└── build/
    ├── Dockerfile                    # Production image
    └── Dockerfile.dev                # Dev image with Air hot-reload
```

**イベント駆動設計**: `main.go` には以下のイベントハンドラを組み込んでいます:
- `MessageHandle` — `CONTROL_TEXT_CH`で`m`を受け取ったら、`MEETING_VC`のmute/unmuteを切り替えます
- `VoiceHandle` — `LOBBY_VC`から`MEETING_VC`に移動する際に自動でmute/unmuteを切り替えます

## 重要な内容

- 日本語のProjectです
- Go module pathは `among_us_assist_bot`
- Git flowに従います: `develop` がmainのbranchです
- Goソースコードはプロジェクトルート直下の`app`ディレクトリにあります
- Discord API内容を認識ください：https://discord.com/developers/docs/reference
- 
