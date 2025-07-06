#!/bin/sh
set -e

# 使い方
# ./makeToDoFileFromPR.sh PR番号

case "$1" in
    ''|*[!0-9]*)
        echo "使い方: $0 PR番号 (数字のみ)" >&2
        exit 1
        ;;
esac

cd "$(dirname "$0")/.."

prompt=$(cat << EOS
プルリクエスト #${1} のレビューコメントをタスク化し tmp/todo 配下に1ファイル1タスクで保存して。
- ファイルフォーマットは docs/templates/todo.md に従うこと。
- レビューコメントの取得はgithub MCPを利用すること。
- リポジトリの情報は git remote -v でで確認すること。
- 解決済みのレビューコメントはタスク化しないこと。
EOS
)
echo "prompt >>>"
echo "$prompt"

echo "run gemini >>>"
gemini -y -m "gemini-2.5-flash" -p "$prompt"
