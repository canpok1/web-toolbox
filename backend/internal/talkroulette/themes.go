package talkroulette

// TalkRouletteTheme は単一のテーマを表します。
type TalkRouletteTheme struct {
	ID    string `json:"id"`    // ID
	Genre string `json:"genre"` // ジャンル
	Theme string `json:"theme"` // テーマ
}

// allThemes はハードコードされた全てのトークルーレットのテーマを保持します。
// UUID は事前に生成され、ハードコードされています。
var allThemes = []TalkRouletteTheme{
	// カテゴリ: all
	{ID: "d2f7bb8a-6c99-4b1f-8f2e-0dc8a0f7c2d5", Genre: "all", Theme: "最近ハマっていることは？"},
	{ID: "1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed", Genre: "all", Theme: "行ってみたい場所はどこですか？"},
	{ID: "f3e9c24a-8f54-48f0-8b0f-09a8b0e3c8f7", Genre: "all", Theme: "あなたの好きな映画は何ですか？"},
	{ID: "a7b3c8d9-9e2f-4a7b-8f0c-1d2e3f4a5b6c", Genre: "all", Theme: "子供の頃の夢は何でしたか？"},
	{ID: "c5d8e9f0-1a2b-4c3d-8e9f-0a1b2c3d4e5f", Genre: "all", Theme: "休日の過ごし方について教えてください。"},
	{ID: "d9e0f1a2-3b4c-5d6e-8f0a-1b2c3d4e5f6a", Genre: "all", Theme: "今までで一番嬉しかったことは？"},
	{ID: "e0f1a2b3-c4d5-6e7f-8a0b-1c2d3e4f5a6b", Genre: "all", Theme: "何か一つ、特技を教えてください。"},
	{ID: "f1a2b3c4-d5e6-7f8a-9b0c-1d2e3f4a5b6d", Genre: "all", Theme: "あなたの好きな食べ物は何ですか？"},
	{ID: "a2b3c4d5-e6f7-8a9b-0c1d-2e3f4a5b6c7e", Genre: "all", Theme: "尊敬する人はいますか？"},
	{ID: "b3c4d5e6-f78a-9b0c-1d2e-3f4a5b6c7d8f", Genre: "all", Theme: "今のあなたを漢字一文字で表すと？"},
	{ID: "c4d5e6f7-8a9b-0c1d-2e3f-4a5b6c7e8f0a", Genre: "all", Theme: "タイムマシンがあったら、過去と未来どっちに行きたい？"},
	{ID: "d5e6f78a-9b0c-1d2e-3f4a-5b6c7d8e9f0b", Genre: "all", Theme: "宝くじが当たったら何をしたい？"},
	{ID: "e6f78a9b-0c1d-2e3f-4a5b-6c7d8e9f0a1c", Genre: "all", Theme: "宇宙人と友達になれると思いますか？"},
	{ID: "f78a9b0c-1d2e-3f4a-5b6c-7d8e9f0a1b2d", Genre: "all", Theme: "もし透明人間になれたら、何をしたい？"},
	{ID: "8a9b0c1d-2e3f-4a5b-6c7d-8e9f0a1b2c3e", Genre: "all", Theme: "10年後の自分にメッセージを送るとしたら、何て書く？"},
	{ID: "9b0c1d2e-3f4a-5b6c-7d8e-9f0a1b2c3d4f", Genre: "all", Theme: "無人島に一つだけ持っていくなら何？"},
	{ID: "0c1d2e3f-4a5b-6c7d-8e9f-0a1b2c3d4e5a", Genre: "all", Theme: "超能力が使えるなら、どんな力が欲しい？"},
	{ID: "1d2e3f4a-5b6c-7d8e-9f0a-1b2c3d4e5f6b", Genre: "all", Theme: "学生時代の一番の思い出は何ですか？"},
	{ID: "2e3f4a5b-6c7d-8e9f-0a1b-2c3d4e5f6a7c", Genre: "all", Theme: "あなたのストレス解消法は？"},
	{ID: "3f4a5b6c-7d8e-9f0a-1b2c-3d4e5f6a7b8d", Genre: "all", Theme: "人生で一番影響を受けた本は？"},
	{ID: "4a5b6c7d-8e9f-0a1b-2c3d-4e5f6a7b8c9e", Genre: "all", Theme: "好きな音楽やアーティストは？"},
	{ID: "5b6c7d8e-9f0a-1b2c-3d4e-5f6a7b8c9d0f", Genre: "all", Theme: "ペットを飼っていますか？どんな動物？"},
	{ID: "6c7d8e9f-0a1b-2c3d-4e5f-6a7b8c9d0e1a", Genre: "all", Theme: "あなたのラッキーカラーは何色ですか？"},
	{ID: "7d8e9f0a-1b2c-3d4e-5f6a-7b8c9d0e1f2b", Genre: "all", Theme: "旅行に行くなら、どんなスタイルが好き？"},
	{ID: "8e9f0a1b-2c3d-4e5f-6a7b-8c9d0e1f2a3c", Genre: "all", Theme: "得意料理を教えてください。"},
	{ID: "9f0a1b2c-3d4e-5f6a-7b8c-9d0e1f2a3b4d", Genre: "all", Theme: "朝型？夜型？"},
	{ID: "0a1b2c3d-4e5f-6a7b-8c9d-0e1f2a3b4c5e", Genre: "all", Theme: "最近感動したことは？"},
	{ID: "1b2c3d4e-5f6a-7b8c-9d0e-1f2a3b4c5d6f", Genre: "all", Theme: "どんな時に幸せを感じますか？"},
	{ID: "2c3d4e5f-6a7b-8c9d-0e1f-2a3b4c5d6e0a", Genre: "all", Theme: "あなたの元気の源は？"},
	{ID: "3d4e5f6a-7b8c-9d0e-1f2a-3b4c5d6e0f1b", Genre: "all", Theme: "座右の銘はありますか？"},

	// カテゴリ: general
	{ID: "4e5f6a7b-8c9d-0e1f-2a3b-4c5d6e0f1a2c", Genre: "general", Theme: "最近、何か面白いことありました？"},
	{ID: "5f6a7b8c-9d0e-1f2a-3b4c-5d6e0f1a2b3d", Genre: "general", Theme: "最近気になっているニュースは？"},
	{ID: "6a7b8c9d-0e1f-2a3b-4c5d-6e0f1a2b3c4e", Genre: "general", Theme: "今の仕事を選んだ理由は？"},
	{ID: "7b8c9d0e-1f2a-3b4c-5d6e-0f1a2b3c4d5f", Genre: "general", Theme: "学生時代に熱中していたことは？"},
	{ID: "8c9d0e1f-2a3b-4c5d-6e0f-1a2b3c4d5e6a", Genre: "general", Theme: "最近買ったお気に入りのものは？"},
	{ID: "9d0e1f2a-3b4c-5d6e-0f1a-2b3c4d5e6f7b", Genre: "general", Theme: "最近挑戦したことは？"},
	{ID: "0e1f2a3b-4c5d-6e0f-1a2b-3c4d5e6f7a8c", Genre: "general", Theme: "最近嬉しかったことは？"},
	{ID: "1f2a3b4c-5d6e-0f1a-2b3c-4d5e6f7a8b9d", Genre: "general", Theme: "最近悲しかったことは？"},
	{ID: "2a3b4c5d-6e0f-1a2b-3c4d-5e6f7a8b9c0e", Genre: "general", Theme: "最近驚いたことは？"},
	{ID: "3b4c5d6e-0f1a-2b3c-4d5e-6f7a8b9c0d1f", Genre: "general", Theme: "最近考えたことは？"},

	// カテゴリ: hobby
	{ID: "4c5d6e0f-1a2b-3c4d-5e6f-7a8b9c0d1e2a", Genre: "hobby", Theme: "あなたの趣味は何ですか？"},
	{ID: "5d6e0f1a-2b3c-4d5e-6f7a-8b9c0d1e2f3b", Genre: "hobby", Theme: "その趣味を始めたきっかけは？"},
	{ID: "6e0f1a2b-3c4d-5e6f-7a8b-9c0d1e2f3a4c", Genre: "hobby", Theme: "趣味の魅力を教えてください。"},
	{ID: "0f1a2b3c-4d5e-6f7a-8b9c-0d1e2f3a4b5d", Genre: "hobby", Theme: "趣味を通して得られたものは？"},
	{ID: "1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6e", Genre: "hobby", Theme: "これから挑戦したい趣味は？"},
	{ID: "2b3c4d5e-6f7a-8b9c-0d1e-2f3a4b5c6d7f", Genre: "hobby", Theme: "趣味にかける時間は？"},
	{ID: "3c4d5e6f-7a8b-9c0d-1e2f-3a4b5c6d7e8a", Genre: "hobby", Theme: "趣味の道具でこだわりのものは？"},
	{ID: "4d5e6f7a-8b9c-0d1e-2f3a-4b5c6d7e8f9b", Genre: "hobby", Theme: "他の人におすすめしたい趣味は？"},
	{ID: "5e6f7a8b-9c0d-1e2f-3a4b-5c6d7e8f9a0c", Genre: "hobby", Theme: "趣味の集まりには参加してる？"},
	{ID: "6f7a8b9c-0d1e-2f3a-4b5c-6d7e8f9a0b1d", Genre: "hobby", Theme: "あなたの自慢の趣味グッズは？"},

	// カテゴリ: food
	{ID: "7a8b9c0d-1e2f-3a4b-5c6d-7e8f9a0b1c2e", Genre: "food", Theme: "好きな食べ物は何ですか？"},
	{ID: "8b9c0d1e-2f3a-4b5c-6d7e-8f9a0b1c2d3f", Genre: "food", Theme: "嫌いな食べ物は何ですか？"},
	{ID: "9c0d1e2f-3a4b-5c6d-7e8f-9a0b1c2d3e4a", Genre: "food", Theme: "得意料理を教えてください。"},
	{ID: "0d1e2f3a-4b5c-6d7e-8f9a-0b1c2d3e4f5b", Genre: "food", Theme: "最近作った料理は？"},
	{ID: "1e2f3a4b-5c6d-7e8f-9a0b-1c2d3e4f5a6c", Genre: "food", Theme: "よく行くお店はありますか？"},
	{ID: "2f3a4b5c-6d7e-8f9a-0b1c-2d3e4f5a6b7d", Genre: "food", Theme: "おすすめのレストランは？"},
	{ID: "3a4b5c6d-7e8f-9a0b-1c2d-3e4f5a6b7c8e", Genre: "food", Theme: "料理をする頻度は？"},
	{ID: "4b5c6d7e-8f9a-0b1c-2d3e-4f5a6b7c8d9f", Genre: "food", Theme: "お弁当は作りますか？"},
	{ID: "5c6d7e8f-9a0b-1c2d-3e4f-5a6b7c8d9e0a", Genre: "food", Theme: "外食する頻度は？"},
	{ID: "6d7e8f9a-0b1c-2d3e-4f5a-6b7c8d9e0f1b", Genre: "food", Theme: "料理でこだわっていることは？"},

	// カテゴリ: travel
	{ID: "7e8f9a0b-1c2d-3e4f-5a6b-7c8d9e0f1a2c", Genre: "travel", Theme: "今までで一番思い出に残っている旅行は？"},
	{ID: "8f9a0b1c-2d3e-4f5a-6b7c-8d9e0f1a2b3d", Genre: "travel", Theme: "一番良かった旅行先は？"},
	{ID: "9a0b1c2d-3e4f-5a6b-7c8d-9e0f1a2b3c4e", Genre: "travel", Theme: "おすすめの旅行先は？"},
	{ID: "a0b1c2d3-e4f5-a6b7-c8d9-e0f1a2b3c4d5", Genre: "travel", Theme: "旅行の計画は自分で立てる？"},
	{ID: "b1c2d3e4-f5a6-b7c8-d9e0-f1a2b3c4d5e6", Genre: "travel", Theme: "旅行に行く頻度は？"},
	{ID: "c2d3e4f5-a6b7-c8d9-e0f1-a2b3c4d5e6f7", Genre: "travel", Theme: "日帰り旅行で行くならどこ？"},
	{ID: "d3e4f5a6-b7c8-d9e0-f1a2-b3c4d5e6f7a8", Genre: "travel", Theme: "国内旅行で行きたい場所は？"},
	{ID: "e4f5a6b7-c8d9-e0f1-a2b3-c4d5e6f7a8b9", Genre: "travel", Theme: "海外旅行で行きたい場所は？"},
	{ID: "f5a6b7c8-d9e0-f1a2-b3c4-d5e6f7a8b9ca", Genre: "travel", Theme: "旅行のスタイルは？"},
	{ID: "a6b7c8d9-e0f1-a2b3-c4d5-e6f7a8b9cabd", Genre: "travel", Theme: "旅行の必需品は？"},
}
