# gkit

Git CLI Tools - Git 操作を簡略化する CLI ツール群

## 概要

gkit は、日常的な Git 操作を短縮化するコマンドラインツールです。各コマンドは独立して動作し、対応する Git コマンドの短縮形として機能します。

## インストール

### GitHub Releases から

```bash
# 最新版をインストール（Linux/macOS）
curl -sL https://github.com/wato787/gkit/releases/latest/download/install.sh | bash

# 手動インストール
# https://github.com/wato787/gkit/releases から該当するOSのバイナリをダウンロード
# /usr/local/bin/ に配置
```

### go install から

```bash
# 一括インストール（推奨）
go install github.com/wato787/gkit/cmd/gkit-installer@latest && gkit-installer

# 個別インストール
go install github.com/wato787/gkit/cmd/gs@latest
go install github.com/wato787/gkit/cmd/ga@latest
go install github.com/wato787/gkit/cmd/gc@latest
go install github.com/wato787/gkit/cmd/gp@latest
```

### ソースからビルド

```bash
git clone https://github.com/wato787/gkit.git
cd gkit
make install
```

## コマンド

### gs (git switch)

```bash
gs main                    # git switch main
gs develop                 # git switch develop
gs -                       # git switch - (前のブランチに戻る)
gs -c feature/new-branch   # git switch -c feature/new-branch
gs -c f/new-branch         # feature/new-branch に展開して作成
```

### ga (git add)

```bash
ga .                       # git add .
ga file.txt                # git add file.txt
ga src/                    # git add src/
ga *.go                    # git add *.go
```

### gc (git commit)

```bash
gc "commit message"        # git commit -m "commit message"
gc                         # エディタでコミットメッセージ入力
```

### gp (git push)

```bash
gp                         # git push
gp origin main             # git push origin main
gp origin feature/branch   # git push origin feature/branch
```

## プレフィックス展開

`gs -c` でブランチ作成時のみ適用される短縮形：

| 短縮形 | 展開後     |
| ------ | ---------- |
| `f/`   | `feature/` |
| `fix/` | `fix/`     |
| `h/`   | `hotfix/`  |
| `r/`   | `release/` |
| `b/`   | `bugfix/`  |
| `e/`   | `epic/`    |

## タブ補完

### bash

```bash
echo 'source <(gs completion bash)' >> ~/.bashrc
echo 'source <(ga completion bash)' >> ~/.bashrc
echo 'source <(gc completion bash)' >> ~/.bashrc
echo 'source <(gp completion bash)' >> ~/.bashrc
```

### zsh

```bash
echo 'source <(gs completion zsh)' >> ~/.zshrc
echo 'source <(ga completion zsh)' >> ~/.zshrc
echo 'source <(gc completion zsh)' >> ~/.zshrc
echo 'source <(gp completion zsh)' >> ~/.zshrc
```

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
