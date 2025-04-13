export interface Round {
  id: string;
  status: RoundStatus;
}

export type RoundStatus = "voting" | "revealed";
