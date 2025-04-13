export interface Round {
  roundId: string;
  sessionId: string;
  status: RoundStatus;
}

export type RoundStatus = "voting" | "revealed";
