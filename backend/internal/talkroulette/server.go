package talkroulette

import (
	"math/rand"
	"time"
)

// init関数はパッケージの初期化時に一度だけ実行されます。
// ここで乱数ジェネレータのシードを設定します。
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetTalkRouletteThemesLogic はテーマ取得のコアロジックを含みます。
func GetTalkRouletteThemesLogic(queryGenre *string, queryMaxCount *int) ([]TalkRouletteTheme, error) {
	// 乱数シードはinit()関数で初期化時に設定されます。

	maxCount := 20 // openapi.yml からのデフォルト値
	if queryMaxCount != nil {
		if *queryMaxCount >= 1 && *queryMaxCount <= 100 {
			maxCount = *queryMaxCount
		}
		// 値が範囲外の場合、デフォルトの20を使用します。
	}

	var candidateThemes []TalkRouletteTheme
	if queryGenre != nil && *queryGenre != "" {
		for _, theme := range allThemes {
			if theme.Genre == *queryGenre {
				candidateThemes = append(candidateThemes, theme)
			}
		}
	} else {
		candidateThemes = make([]TalkRouletteTheme, len(allThemes))
		copy(candidateThemes, allThemes) // グローバルな allThemes をシャッフルしないようにコピーを作成します。
	}

	var selectedThemes []TalkRouletteTheme
	if len(candidateThemes) > 0 {
		if len(candidateThemes) <= maxCount {
			selectedThemes = candidateThemes
		} else {
			// サブセットを選択する必要がある場合はコピーをシャッフルします。
			shuffledCandidates := make([]TalkRouletteTheme, len(candidateThemes))
			copy(shuffledCandidates, candidateThemes)
			rand.Shuffle(len(shuffledCandidates), func(i, j int) {
				shuffledCandidates[i], shuffledCandidates[j] = shuffledCandidates[j], shuffledCandidates[i]
			})
			selectedThemes = shuffledCandidates[:maxCount]
		}
	} else {
		selectedThemes = []TalkRouletteTheme{} // nilではなく空のスライスを返します。
	}

	return selectedThemes, nil
}
