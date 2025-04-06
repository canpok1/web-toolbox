/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
    "/api/planning-poker/ws": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * リアルタイム更新のための WebSocket エンドポイント
         * @description プランニングポーカーセッションに関するリアルタイム更新を受信するための WebSocket 接続を確立します。
         *     サーバーは以下のイベントを含むメッセージを送信します:
         *     - participantJoined: 新しい参加者がセッションに参加しました。
         *     - roundStarted: 新しいラウンドが開始されました。
         *     - voteSubmitted: 参加者が投票を送信しました。
         *     - votesRevealed: ラウンドの投票が公開されました。
         *     - sessionEnded: セッションが終了しました。
         *
         */
        get: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description プロトコルの切り替え */
                101: {
                    headers: {
                        /** @description websocket */
                        Upgrade?: string;
                        /** @description Upgrade */
                        Connection?: string;
                        [name: string]: unknown;
                    };
                    content?: never;
                };
                /** @description 不正なリクエスト */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content?: never;
                };
                /** @description サーバー内部エラー */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content?: never;
                };
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/planning-poker/sessions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** セッションを作成する */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody: {
                content: {
                    "application/json": components["schemas"]["CreateSessionRequest"];
                };
            };
            responses: {
                /** @description セッション作成成功 */
                201: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["CreateSessionResponse"];
                    };
                };
                /** @description 不正なリクエスト */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
                /** @description サーバーエラー */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/planning-poker/sessions/{sessionId}/participants": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** セッションに参加する */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    /** @description セッションID */
                    sessionId: string;
                };
                cookie?: never;
            };
            requestBody: {
                content: {
                    "application/json": components["schemas"]["JoinSessionRequest"];
                };
            };
            responses: {
                /** @description セッション参加成功 */
                201: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["JoinSessionResponse"];
                    };
                };
                /** @description 不正なリクエスト */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
                /** @description セッションが見つからない */
                404: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
                /** @description サーバーエラー */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/planning-poker/sessions/{sessionId}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** セッションを取得する */
        get: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    /** @description セッションID */
                    sessionId: string;
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description セッション取得成功 */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["GetSessionResponse"];
                    };
                };
                /** @description セッションが見つからない */
                404: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
                /** @description サーバーエラー */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/planning-poker/sessions/{sessionId}/end": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** セッションを終了する */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    /** @description セッションID */
                    sessionId: string;
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description セッション終了成功 */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["EndSessionResponse"];
                    };
                };
                /** @description セッションが見つからない */
                404: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
                /** @description サーバーエラー */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/planning-poker/sessions/{sessionId}/rounds": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** ラウンドを開始する */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    /** @description セッションID */
                    sessionId: string;
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description ラウンド開始成功 */
                201: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["StartRoundResponse"];
                    };
                };
                /** @description 不正なリクエスト */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
                /** @description セッションが見つからない */
                404: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
                /** @description サーバーエラー */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/planning-poker/rounds/{roundId}/reveal": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** ラウンドを終了する */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    /** @description ラウンドID */
                    roundId: string;
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description ラウンド終了成功 */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["RevealRoundResponse"];
                    };
                };
                /** @description ラウンドが見つからない */
                404: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
                /** @description サーバーエラー */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/planning-poker/rounds/{roundId}/votes": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** 投票を送信する */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    /** @description ラウンドID */
                    roundId: string;
                };
                cookie?: never;
            };
            requestBody: {
                content: {
                    "application/json": components["schemas"]["SendVoteRequest"];
                };
            };
            responses: {
                /** @description 投票送信成功 */
                201: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["SendVoteResponse"];
                    };
                };
                /** @description 不正なリクエスト */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
                /** @description ラウンドが見つからない */
                404: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
                /** @description サーバーエラー */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
}
export type webhooks = Record<string, never>;
export interface components {
    schemas: {
        /** @description エラーレスポンス */
        ErrorResponse: {
            /** @description エラーメッセージ */
            message: string;
        };
        /**
         * @description スケールの種類
         * @enum {string}
         */
        ScaleType: "fibonacci" | "t-shirt" | "power-of-two" | "custom";
        /** @description セッション作成リクエスト */
        CreateSessionRequest: {
            /** @description セッション名 */
            sessionName: string;
            /** @description スケールの種類 */
            scaleType: components["schemas"]["ScaleType"];
            /** @description カスタムスケール（scaleTypeがcustomの場合のみ有効） */
            customScale?: string[];
            /** @description ホスト名 */
            hostName: string;
        };
        /** @description セッション作成レスポンス */
        CreateSessionResponse: {
            /**
             * Format: uuid
             * @description 作成されたセッションのID
             */
            sessionId: string;
            /**
             * Format: uuid
             * @description セッションのホストID
             */
            hostId: string;
        };
        /** @description セッション参加リクエスト */
        JoinSessionRequest: {
            /** @description 参加者の名前 */
            name: string;
        };
        /** @description セッション参加レスポンス */
        JoinSessionResponse: {
            /**
             * Format: uuid
             * @description 参加者のID
             */
            participantId: string;
        };
        /** @description セッション取得レスポンス */
        GetSessionResponse: {
            /** @description セッション情報 */
            session: components["schemas"]["Session"];
        };
        /** @description セッション終了レスポンス */
        EndSessionResponse: Record<string, never>;
        /** @description ラウンド開始レスポンス */
        StartRoundResponse: {
            /**
             * Format: uuid
             * @description 開始されたラウンドのID
             */
            roundId: string;
        };
        /** @description ラウンド結果公開レスポンス */
        RevealRoundResponse: Record<string, never>;
        /** @description 投票送信リクエスト */
        SendVoteRequest: {
            /**
             * Format: uuid
             * @description 投票者のID
             */
            participantId: string;
            /** @description 投票値 */
            value: string;
        };
        /** @description 投票送信レスポンス */
        SendVoteResponse: {
            /**
             * Format: uuid
             * @description 投票ID
             */
            voteId: string;
        };
        /** @description セッション参加者 */
        SessionParticipant: {
            /** @description 参加者のID */
            participantId: string;
            /** @description 参加者の名前 */
            name: string;
        };
        /** @description セッション情報 */
        Session: {
            /**
             * Format: uuid
             * @description セッションのID
             */
            sessionId: string;
            /** @description セッション名 */
            sessionName: string;
            /**
             * Format: uuid
             * @description セッションのホストID
             */
            hostId: string;
            /** @description スケールの種類 */
            scaleType: components["schemas"]["ScaleType"];
            /** @description カスタムスケール（scaleTypeがcustomの場合のみ有効） */
            customScale: string[];
            /**
             * Format: uuid
             * @description 現在のラウンドID
             */
            currentRoundId?: string;
            /** @description セッションの状態 */
            status: string;
            /** @description 参加者リスト */
            participants: components["schemas"]["SessionParticipant"][];
            /**
             * Format: date-time
             * @description セッションの作成日時
             */
            createdAt: string;
            /**
             * Format: date-time
             * @description セッションの最終更新日時
             */
            updatedAt: string;
        };
        /** @description WebSocketメッセージ */
        WebSocketMessage: {
            /**
             * @description イベント名
             * @enum {string}
             */
            event?: "participantJoined" | "roundStarted" | "voteSubmitted" | "votesRevealed" | "sessionEnded";
            /** @description イベントのペイロード */
            payload?: components["schemas"]["ParticipantJoinedPayload"] | components["schemas"]["RoundStartedPayload"] | components["schemas"]["VoteSubmittedPayload"] | components["schemas"]["VotesRevealedPayload"] | components["schemas"]["SessionEndedPayload"];
        };
        /** @description participantJoinedイベントのペイロード */
        ParticipantJoinedPayload: {
            /**
             * Format: uuid
             * @description 参加者のID
             */
            participantId?: string;
            /** @description 参加者の名前 */
            name?: string;
        };
        /** @description roundStartedイベントのペイロード */
        RoundStartedPayload: {
            /**
             * Format: uuid
             * @description ラウンドのID
             */
            roundId?: string;
        };
        /** @description voteSubmittedイベントのペイロード */
        VoteSubmittedPayload: {
            /**
             * Format: uuid
             * @description 投票した参加者のID
             */
            participantId?: string;
        };
        /** @description 投票情報 */
        Vote: {
            /**
             * Format: uuid
             * @description 参加者のID
             */
            participantId?: string;
            /** @description 投票値 */
            value?: string;
        };
        /** @description votesRevealedイベントのペイロード */
        VotesRevealedPayload: {
            /** @description 投票結果のリスト */
            votes?: components["schemas"]["Vote"][];
            /**
             * Format: float
             * @description 投票値の平均
             */
            average?: number;
            /**
             * Format: float
             * @description 投票値の中央値
             */
            median?: number;
        };
        /** @description sessionEndedイベントのペイロード */
        SessionEndedPayload: Record<string, never>;
    };
    responses: never;
    parameters: never;
    requestBodies: never;
    headers: never;
    pathItems: never;
}
export type $defs = Record<string, never>;
export type operations = Record<string, never>;
