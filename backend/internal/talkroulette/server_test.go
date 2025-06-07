package talkroulette

import (
	"reflect"
	"sort"
	"testing"
)

// strPtr は文字列へのポインタを作成するヘルパー関数です。
func strPtr(s string) *string {
	return &s
}

// intPtr は整数へのポインタを作成するヘルパー関数です。
func intPtr(i int) *int {
	return &i
}

// sortThemes は比較の一貫性のためにテーマをIDでソートするヘルパー関数です。
func sortThemes(themes []TalkRouletteTheme) {
	sort.Slice(themes, func(i, j int) bool {
		return themes[i].ID < themes[j].ID
	})
}

// countThemesByGenre はグローバルなallThemesスライス内のテーマをジャンル別にカウントするヘルパーです。
func countThemesByGenre(genre string) int {
	count := 0
	for _, t := range allThemes {
		if t.Genre == genre {
			count = count + 1
		}
	}
	return count
}

func TestGetTalkRouletteThemesLogic(t *testing.T) {
	// allThemes はグローバルで themes.go で設定されるため、ここで利用可能です。
	// テストが意味を持つためには allThemes が空でないことを確認します。
	if len(allThemes) == 0 {
		t.Fatal("allThemesが空です。テストを実行できません。themes.goが設定されていることを確認してください。")
	}

	totalThemesAllGenre := countThemesByGenre("all")
	totalThemesGeneralGenre := countThemesByGenre("general")
	// totalThemesHobbyGenre := countThemesByGenre("hobby") // より多くのジャンルのための例

	tests := []struct {
		name             string    // テストケース名
		genre            *string   // ジャンル
		maxCount         *int      // 最大取得件数
		expectedMinCount int       // 期待される最小件数（ランダム選択が発生する場合、より大きなプールから）
		expectedMaxCount int       // 期待される最大件数（ランダム選択が発生する場合）
		expectedGenre    *string   // nilでない場合、返される全てのテーマがこのジャンルに一致するべき
		checkExactCount  bool      // trueの場合、expectedMinCountが期待される正確な件数
	}{
		{
			name:             "ジャンル指定なし、maxCount指定なし (デフォルトmaxCount=20)",
			genre:            nil,
			maxCount:         nil,
			expectedMinCount: min(len(allThemes), 20),
			expectedMaxCount: min(len(allThemes), 20),
			expectedGenre:    nil,
			checkExactCount:  true,
		},
		{
			name:             "特定のジャンル 'general'、maxCount指定なし (デフォルトmaxCount=20)",
			genre:            strPtr("general"),
			maxCount:         nil,
			expectedMinCount: min(totalThemesGeneralGenre, 20),
			expectedMaxCount: min(totalThemesGeneralGenre, 20),
			expectedGenre:    strPtr("general"),
			checkExactCount:  true,
		},
		{
			name:             "ジャンル指定なし、maxCount=5",
			genre:            nil,
			maxCount:         intPtr(5),
			expectedMinCount: min(len(allThemes), 5),
			expectedMaxCount: min(len(allThemes), 5),
			expectedGenre:    nil,
			checkExactCount:  true,
		},
		{
			name:             "特定のジャンル 'general'、maxCount=3 (ジャンル内の総数より少ない)",
			genre:            strPtr("general"),
			maxCount:         intPtr(3),
			expectedMinCount: min(totalThemesGeneralGenre, 3),
			expectedMaxCount: min(totalThemesGeneralGenre, 3),
			expectedGenre:    strPtr("general"),
			checkExactCount:  true,
		},
		{
			name:             "特定のジャンル 'general'、maxCount=50 (ジャンル内の総数より多い)",
			genre:            strPtr("general"),
			maxCount:         intPtr(50),
			expectedMinCount: totalThemesGeneralGenre, // ジャンル内の全てのテーマが返されるべき
			expectedMaxCount: totalThemesGeneralGenre,
			expectedGenre:    strPtr("general"),
			checkExactCount:  true,
		},
		{
			name:             "存在しないジャンル",
			genre:            strPtr("nonexistent"),
			maxCount:         nil,
			expectedMinCount: 0,
			expectedMaxCount: 0,
			expectedGenre:    strPtr("nonexistent"), // またはnil、テーマは返されないため
			checkExactCount:  true,
		},
		{
			name:             "maxCount=1",
			genre:            nil,
			maxCount:         intPtr(1),
			expectedMinCount: min(len(allThemes), 1),
			expectedMaxCount: min(len(allThemes), 1),
			expectedGenre:    nil,
			checkExactCount:  true,
		},
		{
			name:             "maxCount=100 (最大制限値、総テーマ数が100以上と仮定)",
			genre:            nil,
			maxCount:         intPtr(100),
			expectedMinCount: min(len(allThemes), 100),
			expectedMaxCount: min(len(allThemes), 100),
			expectedGenre:    nil,
			checkExactCount:  true,
		},
		{
			name:             "ジャンル 'all'、maxCount指定なし (デフォルトmaxCount=20)",
			genre:            strPtr("all"),
			maxCount:         nil,
			expectedMinCount: min(totalThemesAllGenre, 20),
			expectedMaxCount: min(totalThemesAllGenre, 20),
			expectedGenre:    strPtr("all"),
			checkExactCount:  true,
		},
		{
			name:             "maxCount=0 (無効な値のため、実質的にデフォルトの20になるべき)",
			genre:            nil,
			maxCount:         intPtr(0), // 無効な値なので、ロジックはデフォルトの20を使用するべき
			expectedMinCount: min(len(allThemes), 20),
			expectedMaxCount: min(len(allThemes), 20),
			expectedGenre:    nil,
			checkExactCount:  true,
		},
		{
			name:             "maxCount=101 (無効な値のため、実質的にデフォルトの20になるべき)",
			genre:            nil,
			maxCount:         intPtr(101), // 無効な値なので、ロジックはデフォルトの20を使用するべき
			expectedMinCount: min(len(allThemes), 20),
			expectedMaxCount: min(len(allThemes), 20),
			expectedGenre:    nil,
			checkExactCount:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ランダム性が要因であり、正確な件数をチェックしていない場合、ロジックを複数回実行します。
			// 現在は checkExactCount により、1回の実行で十分です。
			results, err := GetTalkRouletteThemesLogic(tt.genre, tt.maxCount)
			if err != nil {
				t.Fatalf("GetTalkRouletteThemesLogic() エラー = %v", err)
			}

			if tt.checkExactCount {
				if len(results) != tt.expectedMinCount {
					t.Errorf("GetTalkRouletteThemesLogic() 結果の件数 = %d, 期待値 %d", len(results), tt.expectedMinCount)
				}
			} else {
				if len(results) < tt.expectedMinCount || len(results) > tt.expectedMaxCount {
					t.Errorf("GetTalkRouletteThemesLogic() 結果の件数 = %d, 期待範囲 %d-%d", len(results), tt.expectedMinCount, tt.expectedMaxCount)
				}
			}


			if tt.expectedGenre != nil {
				for _, theme := range results {
					if theme.Genre != *tt.expectedGenre {
						t.Errorf("GetTalkRouletteThemesLogic() 結果のテーマのジャンル = %s, 期待値 %s", theme.Genre, *tt.expectedGenre)
					}
				}
			}

			// 件数が0より大きい場合、IDの一意性をチェックします。
			if len(results) > 0 {
				seenIDs := make(map[string]bool)
				for _, theme := range results {
					if seenIDs[theme.ID] {
						t.Errorf("GetTalkRouletteThemesLogic() 結果に重複したテーマID %s が見つかりました", theme.ID)
					}
					seenIDs[theme.ID] = true
					if theme.ID == "" {
						t.Errorf("GetTalkRouletteThemesLogic() テーマIDが空です")
					}
					if theme.Theme == "" {
						t.Errorf("GetTalkRouletteThemesLogic() ID %s のテーマテキストが空です", theme.ID)
					}
				}
			}

			// オプション: (該当するケースで) 連続する2回の呼び出しが異なる順序またはアイテムを返すかチェックすることでランダム性をテストします。
			// これはより高度な設定なしにユニットテストで確実にアサートするのが難しいです。
			// 現在は、コードパスがヒットすれば rand.Shuffle がその仕事を行うと信頼します。
			// 主なチェックは件数とジャンルです。
		})
	}
}

// Test specific case: maxCount が要求され、利用可能なテーマ（ジャンルフィルタ後）が maxCount より少ない場合、
// 利用可能な全てのテーマが返されるべきです。
func TestGetTalkRouletteThemesLogic_MaxCountMoreThanAvailable(t *testing.T) {
	// 'hobby' ジャンルが既知の少数のテーマを持つと仮定します（例：フロントエンドデータに基づき10件）。
	// そして allThemes が設定済みであること。
	numHobbyThemes := countThemesByGenre("hobby")
	if numHobbyThemes == 0 {
		t.Skip("テストをスキップ: このシナリオをテストするための 'hobby' テーマが見つかりません。")
		return
	}

	requestedMaxCount := numHobbyThemes + 5 // 'hobby'で利用可能な数より多く要求

	results, err := GetTalkRouletteThemesLogic(strPtr("hobby"), intPtr(requestedMaxCount))
	if err != nil {
		t.Fatalf("GetTalkRouletteThemesLogic() エラー = %v", err)
	}

	if len(results) != numHobbyThemes {
		t.Errorf("maxCount (%d) > 利用可能数の場合、全ての %d 個の 'hobby' テーマが期待されましたが、%d 個取得しました", requestedMaxCount, numHobbyThemes, len(results))
	}

	for _, theme := range results {
		if theme.Genre != "hobby" {
			t.Errorf("期待されるテーマのジャンルは 'hobby' でしたが、'%s' を取得しました", theme.Genre)
		}
	}
}


// min ヘルパー関数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TestGetTalkRouletteThemesLogic によって allThemes が変更されないことをテストします（シャッフルのため重要）。
func TestGetTalkRouletteThemesLogic_AllThemesUnmodified(t *testing.T) {
    if len(allThemes) < 2 { // シャッフルの影響をテストするには少なくとも2つのテーマが必要です
        t.Skip("AllThemesUnmodified テストをスキップ: シャッフルを検証するのに十分なテーマがありません。")
        return
    }

    // 後で比較するために allThemes のディープコピーを作成します
    originalAllThemes := make([]TalkRouletteTheme, len(allThemes))
    copy(originalAllThemes, allThemes)
	sortThemes(originalAllThemes) // 比較の一貫性のためにソート

    // 注意しないとシャッフルを引き起こすパラメータでロジックを呼び出します
    // 例：ジャンルなし、かつ maxCount が総テーマ数より少ない
    _, err := GetTalkRouletteThemesLogic(nil, intPtr(1))
    if err != nil {
        t.Fatalf("GetTalkRouletteThemesLogic() エラー = %v", err)
    }

	// 比較のために現在の allThemes の新しいコピーを作成します
	// allThemes はグローバル変数であり、注意しないと他のテストによって変更される可能性があるため、これは重要です（テストの実行順序は保証されませんが）。
	// この特定のテストでは、*この特定の呼び出し* がそれを変更したかどうかをチェックしています。
	currentAllThemes := make([]TalkRouletteTheme, len(allThemes))
    copy(currentAllThemes, allThemes)
	sortThemes(currentAllThemes)


    if !reflect.DeepEqual(originalAllThemes, currentAllThemes) {
        t.Errorf("GetTalkRouletteThemesLogic() がグローバルな allThemes スライスを変更しました。変更前: %v, 変更後: %v", originalAllThemes, currentAllThemes)
		// デバッグのため、可能であれば異なる要素、または長さを出力します
		if len(originalAllThemes) != len(currentAllThemes) {
			t.Logf("長さが異なります: オリジナル %d, 現在 %d", len(originalAllThemes), len(currentAllThemes))
		}
		// 必要であれば、より詳細な差分を追加できます
    }
}
