const TalkRouletteTopPage = () => {
  const htmlContent = `<!DOCTYPE html>
<html lang="ja">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ラントークテーマ</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600&display=swap" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css"
        integrity="sha512-9usAa10IRO0HhonpyAIVpjrylPvoDwiPUiKdWk5t3PyolY1cOd4DSE0Ga+ri4AuTroPR5aQvXU9xC6qOPnzFeg=="
        crossorigin="anonymous" referrerpolicy="no-referrer" />
    <style>
        body {
            font-family: 'Inter', sans-serif;
        }

        .feedback-button {
            transition: transform 0.2s ease-in-out, color 0.2s ease-in-out;
        }

        .feedback-button:hover {
            transform: scale(1.2);
        }

        .feedback-button.liked {
            color: #10b981;
        }

        .feedback-button.disliked {
            color: #ef4444;
        }
    </style>
</head>

<body class="bg-gradient-to-r from-gray-100 to-gray-300 flex justify-center items-center min-h-screen p-4">
    <div class="bg-white rounded-lg shadow-xl p-8 max-w-md w-full text-center">
        <h1 class="text-2xl font-semibold text-gray-700 mb-6">今日のトークテーマ</h1>

        <div id="theme-display"
            class="text-xl text-blue-800 font-semibold rounded-md p-4 bg-blue-50/50 mb-6 min-h-[3rem] flex items-center justify-center transition-all duration-300">
            テーマを読み込み中...
        </div>

        <div class="mt-6 space-y-4">
            <div class="flex justify-center items-center gap-8">
                <button id="like-button" class="feedback-button text-gray-600 hover:text-green-500" aria-label="良いテーマ">
                    <i class="fas fa-thumbs-up text-2xl"></i>
                </button>
                <button id="dislike-button" class="feedback-button text-gray-600 hover:text-red-500" aria-label="悪いテーマ">
                    <i class="fas fa-thumbs-down text-2xl"></i>
                </button>
            </div>
            <p id="feedback-message" class="text-gray-500 text-sm"></p>

            <div>
                <label for="genre-select" class="block text-gray-700 text-sm font-bold mb-2">ジャンルを選択:</label>
                <select id="genre-select"
                    class="shadow appearance-none border rounded w-full py-3 px-4 text-gray-700 leading-tight focus:outline-none focus:shadow-outline">
                    <option value="all">すべて</option>
                    <option value="general">一般</option>
                    <option value="hobby">趣味</option>
                    <option value="food">食べ物</option>
                    <option value="travel">旅行</option>
                </select>
            </div>

            <button id="new-theme-button"
                class="bg-gradient-to-r from-blue-600 to-blue-700 hover:from-blue-700 hover:to-blue-800 text-white font-semibold py-3 px-6 rounded-full shadow-md transition duration-300 ease-in-out hover:scale-105">
                別のテーマを引く
            </button>

            <a href="https://forms.gle/your-google-form-url" target="_blank" rel="noopener noreferrer"
                class="text-blue-600 hover:text-blue-800 transition-colors duration-200 text-sm font-semibold block mt-2">
                新しいトークテーマを投稿する
            </a>
            <p class="text-gray-500 text-xs mt-1">
                良いテーマを思いついたら、ぜひ投稿してください！
            </p>
        </div>
    </div>

    <script>
        const themes = {
            "all": [
                "最近ハマっていることは？",
                "行ってみたい場所はどこですか？",
                "あなたの好きな映画は何ですか？",
                "子供の頃の夢は何でしたか？",
                "休日の過ごし方について教えてください。",
                "今までで一番嬉しかったことは？",
                "何か一つ、特技を教えてください。",
                "あなたの好きな食べ物は何ですか？",
                "尊敬する人はいますか？",
                "今のあなたを漢字一文字で表すと？",
                "タイムマシンがあったら、過去と未来どっちに行きたい？",
                "宝くじが当たったら何をしたい？",
                "宇宙人と友達になれると思いますか？",
                "もし透明人間になれたら、何をしたい？",
                "10年後の自分にメッセージを送るとしたら、何て書く？",
                "無人島に一つだけ持っていくなら何？",
                "超能力が使えるなら、どんな力が欲しい？",
                "学生時代の一番の思い出は何ですか？",
                "あなたのストレス解消法は？",
                "人生で一番影響を受けた本は？",
                "好きな音楽やアーティストは？",
                "ペットを飼っていますか？どんな動物？",
                "あなたのラッキーカラーは何色ですか？",
                "旅行に行くなら、どんなスタイルが好き？",
                "得意料理を教えてください。",
                "朝型？夜型？",
                "最近感動したことは？",
                "どんな時に幸せを感じますか？",
                "あなたの元気の源は？",
                "座右の銘はありますか？"
            ],
            "general": [
                "最近、何か面白いことありました？",
                "最近気になっているニュースは？",
                "今の仕事を選んだ理由は？",
                "学生時代に熱中していたことは？",
                "最近買ったお気に入りのものは？",
                "最近挑戦したことは？",
                "最近嬉しかったことは？",
                "最近悲しかったことは？",
                "最近驚いたことは？",
                "最近考えたことは？"
            ],
            "hobby": [
                "あなたの趣味は何ですか？",
                "その趣味を始めたきっかけは？",
                "趣味の魅力を教えてください。",
                "趣味を通して得られたものは？",
                "これから挑戦したい趣味は？",
                "趣味にかける時間は？",
                "趣味の道具でこだわりのものは？",
                "他の人におすすめしたい趣味は？",
                "趣味の集まりには参加してる？",
                "あなたの自慢の趣味グッズは？"
            ],
            "food": [
                "好きな食べ物は何ですか？",
                "嫌いな食べ物は何ですか？",
                "得意料理を教えてください。",
                "最近作った料理は？",
                "よく行くお店はありますか？",
                "おすすめのレストランは？",
                "料理をする頻度は？",
                "お弁当は作りますか？",
                "外食する頻度は？",
                "料理でこだわっていることは？"
            ],
            "travel": [
                "今までで一番思い出に残っている旅行は？",
                "一番良かった旅行先は？",
                "おすすめの旅行先は？",
                "旅行の計画は自分で立てる？",
                "旅行に行く頻度は？",
                "日帰り旅行で行くならどこ？",
                "国内旅行で行きたい場所は？",
                "海外旅行で行きたい場所は？",
                "旅行のスタイルは？",
                "旅行の必需品は？"
            ]
        };

        const likeButton = document.getElementById("like-button");
        const dislikeButton = document.getElementById("dislike-button");
        const feedbackMessage = document.getElementById("feedback-message");
        let liked = false;
        let disliked = false;

        function getRandomTheme(genre) {
            const selectedGenre = genre || "all";
            const genreThemes = themes[selectedGenre] || themes["all"];
            const randomIndex = Math.floor(Math.random() * genreThemes.length);
            return genreThemes[randomIndex];
        }

        function displayTheme() {
            const themeDisplay = document.getElementById("theme-display");
            const genreSelect = document.getElementById("genre-select");
            const selectedGenre = genreSelect.value;
            const randomTheme = getRandomTheme(selectedGenre);
            themeDisplay.textContent = randomTheme;
            // アニメーションのため、一度不透明にしてから表示
            themeDisplay.style.opacity = '0';
            setTimeout(() => {
                themeDisplay.style.opacity = '1';
            }, 100);
            // リセット
            liked = false;
            disliked = false;
            likeButton.classList.remove('liked');
            dislikeButton.classList.remove('disliked');
            feedbackMessage.textContent = '';
        }

        function handleLikeClick() {
            if (!liked) {
                liked = true;
                disliked = false;
                likeButton.classList.add('liked');
                dislikeButton.classList.remove('disliked');
                feedbackMessage.textContent = '良いテーマですね！';
            } else {
                liked = false;
                likeButton.classList.remove('liked');
                feedbackMessage.textContent = '';
            }
        }

        function handleDislikeClick() {
            if (!disliked) {
                disliked = true;
                liked = false;
                dislikeButton.classList.add('disliked');
                likeButton.classList.remove('liked');
                feedbackMessage.textContent = 'テーマを変更しますね。';
            } else {
                disliked = false;
                dislikeButton.classList.remove('disliked');
                feedbackMessage.textContent = '';
            }
        }

        likeButton.addEventListener("click", handleLikeClick);
        dislikeButton.addEventListener("click", handleDislikeClick);

        const newThemeButton = document.getElementById("new-theme-button");
        newThemeButton.addEventListener("click", displayTheme);

        displayTheme();
    </script>
</body>

</html>`;

  return <div dangerouslySetInnerHTML={{ __html: htmlContent }} />;
};

export default TalkRouletteTopPage;
