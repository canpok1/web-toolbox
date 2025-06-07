import createClient from "openapi-fetch";
import type { paths } from "./types/api.gen";

export class TalkRouletteClient {
  readonly client;

  constructor() {
    this.client = createClient<paths>({
      baseUrl: "/",
    });
  }

  async getThemes(
    genre?: string,
    maxCount = 20,
  ): Promise<
    paths["/api/talk-roulette/themes"]["get"]["responses"][200]["content"]["application/json"]
  > {
    const { data, error } = await this.client.GET("/api/talk-roulette/themes", {
      params: {
        query: {
          genre,
          maxCount,
        },
      },
    });
    if (error) {
      throw new Error(error.message);
    }
    return data;
  }
}
