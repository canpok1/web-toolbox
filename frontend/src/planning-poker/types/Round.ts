import type { RoundSummary } from "./RoundSummary";
import type { Vote } from "./Vote";

export interface Round {
  roundId: string;
  sessionId: string;
  status: RoundStatus;
  votes: Vote[];
  summary?: RoundSummary;
}

export type RoundStatus = "voting" | "revealed";
