import createClient from "openapi-fetch";
import type { paths } from "./types/api.gen";

export class ApiClient {
  readonly client;

  constructor() {
    this.client = createClient<paths>({
      baseUrl: "/",
    });
  }

  async createSession(
    body: paths["/api/planning-poker/sessions"]["post"]["requestBody"]["content"]["application/json"],
  ): Promise<
    paths["/api/planning-poker/sessions"]["post"]["responses"][201]["content"]["application/json"]
  > {
    const { data, error } = await this.client.POST(
      "/api/planning-poker/sessions",
      {
        body,
      },
    );
    if (error) {
      throw new Error(error.message);
    }
    return data;
  }

  async joinSession(
    sessionId: string,
    body: paths["/api/planning-poker/sessions/{sessionId}/participants"]["post"]["requestBody"]["content"]["application/json"],
  ): Promise<
    paths["/api/planning-poker/sessions/{sessionId}/participants"]["post"]["responses"][201]["content"]["application/json"]
  > {
    const { data, error } = await this.client.POST(
      "/api/planning-poker/sessions/{sessionId}/participants",
      {
        params: {
          path: {
            sessionId,
          },
        },
        body,
      },
    );
    if (error) {
      throw new Error(error.message);
    }
    return data;
  }

  async fetchSession(
    sessionId: string,
  ): Promise<
    paths["/api/planning-poker/sessions/{sessionId}"]["get"]["responses"][200]["content"]["application/json"]
  > {
    const { data, error } = await this.client.GET(
      "/api/planning-poker/sessions/{sessionId}",
      {
        params: {
          path: {
            sessionId,
          },
        },
      },
    );
    if (error) {
      throw new Error(error.message);
    }
    return data;
  }

  async endSession(
    sessionId: string,
  ): Promise<
    paths["/api/planning-poker/sessions/{sessionId}/end"]["post"]["responses"][200]["content"]["application/json"]
  > {
    const { data, error } = await this.client.POST(
      "/api/planning-poker/sessions/{sessionId}/end",
      {
        params: {
          path: {
            sessionId,
          },
        },
      },
    );
    if (error) {
      throw new Error(error.message);
    }
    return data;
  }

  async fetchRound(
    roundId: string,
  ): Promise<
    paths["/api/planning-poker/rounds/{roundId}"]["get"]["responses"][200]["content"]["application/json"]
  > {
    const { data, error } = await this.client.GET(
      "/api/planning-poker/rounds/{roundId}",
      {
        params: {
          path: {
            roundId,
          },
        },
      },
    );
    if (error) {
      throw new Error(error.message);
    }
    return data;
  }

  async startRound(
    sessionId: string,
  ): Promise<
    paths["/api/planning-poker/sessions/{sessionId}/rounds"]["post"]["responses"][201]["content"]["application/json"]
  > {
    const { data, error } = await this.client.POST(
      "/api/planning-poker/sessions/{sessionId}/rounds",
      {
        params: {
          path: {
            sessionId,
          },
        },
      },
    );
    if (error) {
      throw new Error(error.message);
    }
    return data;
  }

  async revealRound(
    roundId: string,
  ): Promise<
    paths["/api/planning-poker/rounds/{roundId}/reveal"]["post"]["responses"][200]["content"]["application/json"]
  > {
    const { data, error } = await this.client.POST(
      "/api/planning-poker/rounds/{roundId}/reveal",
      {
        params: {
          path: {
            roundId,
          },
        },
      },
    );
    if (error) {
      throw new Error(error.message);
    }
    return data;
  }

  async vote(
    roundId: string,
    body: paths["/api/planning-poker/rounds/{roundId}/votes"]["post"]["requestBody"]["content"]["application/json"],
  ): Promise<
    paths["/api/planning-poker/rounds/{roundId}/votes"]["post"]["responses"][201]["content"]["application/json"]
  > {
    const { data, error } = await this.client.POST(
      "/api/planning-poker/rounds/{roundId}/votes",
      {
        params: {
          path: {
            roundId,
          },
        },
        body,
      },
    );
    if (error) {
      throw new Error(error.message);
    }
    return data;
  }
}
