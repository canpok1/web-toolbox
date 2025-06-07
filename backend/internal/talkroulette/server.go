package talkroulette

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetTalkRouletteThemesLogic はテーマ取得のコアロジックを含みます。
func GetTalkRouletteThemesLogic(queryGenre *string, queryMaxCount *int) ([]TalkRouletteTheme, error) {
	// 通常、randのシードはアプリケーション起動時に一度だけ行うのが望ましいです。
	// しかし、この特定のハンドラ構造では、ここでシードすることにより、このロジックが呼び出された際に確実にシードされることを保証します。
	// この関数がメソッドを持つ構造体の一部であれば、構造体の初期化時にシードを行うことができます。
	rand.Seed(time.Now().UnixNano())

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

// GetTalkRouletteThemesHandler はトークルーレットテーマのHTTPリクエストを処理します。
func GetTalkRouletteThemesHandler(w http.ResponseWriter, r *http.Request) {
	var queryGenrePtr *string
	if genreVal := r.URL.Query().Get("genre"); genreVal != "" {
		queryGenrePtr = &genreVal
	}

	var queryMaxCountPtr *int
	if maxCountValStr := r.URL.Query().Get("maxCount"); maxCountValStr != "" {
		parsedMaxCount, err := strconv.Atoi(maxCountValStr)
		if err == nil { // パースが成功した場合のみ使用します。
			queryMaxCountPtr = &parsedMaxCount
		}
	}

	selectedThemes, errLogic := GetTalkRouletteThemesLogic(queryGenrePtr, queryMaxCountPtr)
	if errLogic != nil {
		// このサンプルロジックはまだエラーを生成しませんが、もしエラーが発生した場合：
		http.Error(w, "テーマの取得に失敗しました: "+errLogic.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Themes []TalkRouletteTheme `json:"themes"`
	}{
		Themes: selectedThemes,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "テーマのJSONへのエンコードに失敗しました", http.StatusInternalServerError)
		return
	}
}
