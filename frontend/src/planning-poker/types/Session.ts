import type { SessionParticipant } from "./Participant";

export interface Session {
  sessionId: string;
  participants: SessionParticipant[];
  currentRoundId?: string;
  hostId: string;
  status: SessionStatus;
}

export const SessionStatusValues = [
  "waiting",
  "inProgress",
  "finished",
] as const;
export type SessionStatus = (typeof SessionStatusValues)[number];

export function isSessionStatus(value: string): value is SessionStatus {
  return SessionStatusValues.some((v) => v === value);
}
