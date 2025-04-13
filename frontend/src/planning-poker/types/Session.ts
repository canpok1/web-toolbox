import type { SessionParticipant } from "./Participant";

export interface Session {
  id: string;
  name: string;
  participants: SessionParticipant[];
  currentRoundId: string | undefined;
  hostId: string;
}
