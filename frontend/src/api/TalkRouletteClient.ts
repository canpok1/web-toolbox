import createClient from "openapi-fetch";
import type { paths } from "./types/api.gen";

// 型エイリアスの定義
type GetThemesResponse =
  paths["/api/talk-roulette/themes"]["get"]["responses"][200]["content"]["application/json"];

export class TalkRouletteClient {
  readonly client;

  constructor() {
    this.client = createClient<paths>({
      baseUrl: "/",
    });
  }

  async getThemes(genre?: string, maxCount = 20): Promise<GetThemesResponse> {
    const { data, error } = await this.client.GET("/api/talk-roulette/themes", {
      params: {
        query: {
          genre,
          maxCount,
        },
      },
    });
    if (error) {
      throw error;
    }
    return data;
  }
}
