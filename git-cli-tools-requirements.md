# Git CLI Tools 要件定義

## 概要
Git操作を簡略化するCLIツール群を作成する。
各コマンドは独立しており、対応するGitコマンドの短縮形として機能する。

## 作成するコマンド

### 1. gs (git switch)
**基本機能**
```bash
gs main                    # git switch main
gs develop                 # git switch develop
gs -                      # git switch - (前のブランチに戻る)
```

**ブランチ作成機能**
```bash
gs -c feature/new-branch   # git switch -c feature/new-branch
gs -c f/new-branch        # feature/new-branch に展開して作成
```

### 2. ga (git add)
**基本機能**
```bash
ga .                      # git add .
ga file.txt               # git add file.txt
ga src/                   # git add src/
ga *.go                   # git add *.go
```

### 3. gc (git commit)
**基本機能**
```bash
gc "commit message"       # git commit -m "commit message"
gc                        # エディタでコミットメッセージ入力
```

### 4. gp (git push)
**基本機能**
```bash
gp                        # git push
gp origin main            # git push origin main
gp origin feature/branch  # git push origin feature/branch
```

## プレフィックス展開ルール

`gs -c` でブランチ作成時のみ適用される短縮形：

| 短縮形 | 展開後 |
|--------|--------|
| `f/`   | `feature/` |
| `h/`   | `hotfix/` |
| `r/`   | `release/` |
| `b/`   | `bugfix/` |

**例**
```bash
gs -c f/user-auth         # → gs -c feature/user-auth
gs -c h/critical-fix      # → gs -c hotfix/critical-fix
gs -c r/v1.2.0           # → gs -c release/v1.2.0
gs -c b/login-error      # → gs -c bugfix/login-error
```

## エラーハンドリング

### 共通
- Gitリポジトリ外で実行された場合の警告
- 不正な引数の場合のヘルプ表示

### gs固有
- 存在しないブランチへの切り替え時の警告
- 既存ブランチと同名のブランチ作成時の警告
- 未コミットの変更がある場合の確認

### ga固有
- 存在しないファイル・ディレクトリの指定時の警告

### gc固有
- ステージングされたファイルがない場合の警告

### gp固有
- リモートリポジトリが設定されていない場合の警告
- プッシュ権限がない場合のエラーハンドリング

## 技術仕様

### 開発言語
- Go

### 使用ライブラリ
- cobra（CLI構築用）
- その他必要に応じて

### 配布方法
- 初期：直接ビルドして個人使用
- 将来：GitHub Releases / go install

## 開発段階

### Phase 1
基本的な4つのコマンド（gs, ga, gc, gp）の実装

### Phase 2
プレフィックス展開機能の実装

### Phase 3
エラーハンドリングの充実

### Phase 4
設定ファイル対応（カスタムプレフィックス等）

## 使用例

```bash
# 新機能開発のワークフロー
gs -c f/user-registration     # feature/user-registration ブランチ作成・切り替え
ga .                         # 変更をステージング
gc "Add user registration"   # コミット
gp                          # プッシュ

# ホットフィックスのワークフロー
gs -c h/security-patch       # hotfix/security-patch ブランチ作成・切り替え
ga src/auth.go              # 特定ファイルをステージング
gc "Fix security vulnerability"  # コミット
gp origin hotfix/security-patch  # 特定ブランチにプッシュ
```
