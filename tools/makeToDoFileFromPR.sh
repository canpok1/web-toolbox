#!/bin/sh

# 使い方
# ./makeToDoFileFromPR.sh PR番号

if [ -z "$1" ]; then
  exit 0
fi

cd "$(dirname "$0")/.."

prompt="PR ${1} のレビューコメントをタスク化して tmp/todo 配下に1ファイル1タスクで保存して。なおファイルのフォーマットは docs/templates/todo.md に従うこと。"
echo "prompt >>> $prompt"

gemini -y -p "$prompt"
