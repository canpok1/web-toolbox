import type { Vote } from "./Vote";

export interface Round {
  roundId: string;
  sessionId: string;
  status: RoundStatus;
  votes: Vote[];
}

export type RoundStatus = "voting" | "revealed";
