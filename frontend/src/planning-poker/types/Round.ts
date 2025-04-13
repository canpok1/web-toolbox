import type { RoundParticipant } from "./Participant";
import type { Vote } from "./Vote";

export interface Round {
  roundId: string;
  sessionId: string;
  status: RoundStatus;
  participants: RoundParticipant[];
  votes: Vote[];
}

export type RoundStatus = "voting" | "revealed";
